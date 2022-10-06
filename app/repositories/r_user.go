package repositories

import (
	"github.com/jinzhu/copier"
	"github.com/shasw94/projX/app/interfaces"
	"github.com/shasw94/projX/app/models"
	"github.com/shasw94/projX/app/models/pivot"
	"github.com/shasw94/projX/app/schema"
	"github.com/shasw94/projX/pkg/errors"
	"github.com/shasw94/projX/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UserRepo user repository struct
type UserRepo struct {
	db interfaces.IDatabase
}

// NewUserRepository return new IUserRepository interface
func NewUserRepository(db interfaces.IDatabase) interfaces.IUserRepository {
	return &UserRepo{db: db}
}

func (u *UserRepo) Register(item *schema.RegisterBodyParams) (*models.User, error) {
	var user models.User
	copier.Copy(&user, &item)
	if err := u.db.GetInstance().Model(&models.User{}).Create(&user).Error; err != nil {
		return nil, errors.ErrorDatabaseCreate.Newm(err.Error())
	}
	return &user, nil
}

func (u *UserRepo) Create(user *models.User) error {
	if err := u.db.GetInstance().Model(&models.User{}).Create(&user).Error; err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}
	return nil
}

func (u *UserRepo) GetByID(id string) (*models.User, error) {
	user := models.User{}
	if err := u.db.GetInstance().Model(&models.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return &user, nil
}

func (u *UserRepo) GetUserByToken(token string) (*models.User, error) {
	var user models.User
	if err := u.db.GetInstance().Model(&models.User{}).Where("refresh_token = ? ", token).First(&user).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return &user, nil
}

func (u *UserRepo) List(param *schema.UserQueryParam) (*[]models.User, error) {
	var query map[string]interface{}
	if err := utils.Copy(&query, &param); err != nil {
		return nil, errors.ErrorMarshal.Newm(err.Error())
	}

	var user []models.User
	if err := u.db.GetInstance().Model(&models.User{}).Where(query).Offset(param.Offset).Limit(param.Limit).Find(&user).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return &user, nil
}

func (u *UserRepo) Login(item *schema.LoginBodyParams) (*models.User, error) {
	user := &models.User{}
	if err := u.db.GetInstance().Model(&models.User{}).Where("username = ?", item.Username).First(&user).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(item.Password))
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return nil, errors.ErrorInvalidPassword.Newm("invalid password")
	}

	return user, nil
}

func (u *UserRepo) RemoveToken(userID string) (*models.User, error) {
	var body = map[string]interface{}{"refresh_token": ""}
	var change models.User
	if err := u.db.GetInstance().Model(&change).Where("id = ?", userID).Updates(body).Error; err != nil {
		return nil, errors.ErrorDatabaseUpdate.Newm(err.Error())
	}
	return &change, nil
}

func (u *UserRepo) Update(userID string, bodyParam *schema.UserUpdateBodyParam) (*models.User, error) {
	var body map[string]interface{}
	err := utils.Copy(&body, &bodyParam)
	if err != nil {
		return nil, errors.ErrorMarshal.Newm(err.Error())
	}

	var change models.User
	if err := u.db.GetInstance().Model(&change).Where("id = ?", userID).Updates(body).Error; err != nil {
		return nil, errors.ErrorDatabaseUpdate.Newm(err.Error())
	}

	return &change, nil
}

// AddPermissions and direct permission to user
// @param string
// @param schema.Permission
// @return error
func (u *UserRepo) AddPermissions(userID string, permissions schema.Permission) error {
	var userPermissions []pivot.UserPermission
	for _, permission := range permissions.Origin() {
		userPermissions = append(userPermissions, pivot.UserPermission{
			UserID:       userID,
			PermissionID: permission.ID,
		})
	}
	return u.db.GetInstance().Clauses(clause.OnConflict{DoNothing: true}).Create(&userPermissions).Error
}

func (u *UserRepo) ReplacePermissions(userID string, permissions schema.Permission) error {
	return u.db.GetInstance().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_permissions.user_id = ?", userID).Delete(&pivot.UserPermission{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		var userPermissions []pivot.UserPermission
		for _, permission := range permissions.Origin() {
			userPermissions = append(userPermissions, pivot.UserPermission{
				UserID:       userID,
				PermissionID: permission.ID,
			})
		}

		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&userPermissions).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})
}

func (u *UserRepo) RemovePermissions(userId string, permissions schema.Permission) error {
	var userPermissions []pivot.UserPermission
	for _, permission := range permissions.Origin() {
		userPermissions = append(userPermissions, pivot.UserPermission{
			UserID:       userId,
			PermissionID: permission.ID,
		})
	}
	return u.db.GetInstance().Delete(&userPermissions).Error
}

func (u *UserRepo) ClearPermissions(userID string) (err error) {
	return u.db.GetInstance().Where("user_permissions.user_id = ?", userID).Delete(&pivot.UserPermission{}).Error
}

func (u *UserRepo) AddRoles(userId string, roles schema.Roles) error {
	var userRoles []pivot.UserRole
	for _, role := range roles.Origin() {
		userRoles = append(userRoles, pivot.UserRole{
			UserID: userId,
			RoleID: role.ID,
		})
	}
	return u.db.GetInstance().Clauses(clause.OnConflict{DoNothing: true}).Create(&userRoles).Error
}

func (u *UserRepo) ReplaceRoles(userId string, roles schema.Roles) error {
	return u.db.GetInstance().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_roles.user_id = ?", userId).Delete(&pivot.UserRole{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		var userRoles []pivot.UserRole
		for _, role := range roles.Origin() {
			userRoles = append(userRoles, pivot.UserRole{
				UserID: userId,
				RoleID: role.ID,
			})
		}
		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&userRoles).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

func (u *UserRepo) RemoveRoles(userId string, roles schema.Roles) error {
	var userRoles []pivot.UserRole
	for _, role := range roles.Origin() {
		userRoles = append(userRoles, pivot.UserRole{
			UserID: userId,
			RoleID: role.ID,
		})
	}
	return u.db.GetInstance().Delete(&userRoles).Error
}

func (u *UserRepo) ClearRoles(userId string) (err error) {
	return u.db.GetInstance().Where("user_roles.user_id = ?", userId).Delete(&pivot.UserRole{}).Error
}

func (u *UserRepo) HasRole(userId string, role models.Role) (b bool, err error) {
	var count int64
	err = u.db.GetInstance().Table("user_roles").Where("user_roles.user_id = ?", userId).Where("user_roles.role_id = ?", role.ID).Count(&count).Error
	return count > 0, err
}

func (u *UserRepo) HasAllRoles(userId string, roles schema.Roles) (b bool, err error) {
	var count int64
	err = u.db.GetInstance().Table("user_roles").Where("user_roles.user_id = ?", userId).Where("user_roles.role_id IN (?)", roles.IDs()).Count(&count).Error
	return roles.Len() == count, err
}

func (u *UserRepo) HasAnyRoles(userID string, roles schema.Roles) (b bool, err error) {
	var count int64
	err = u.db.GetInstance().Table("user_roles").Where("user_roles.user_id = ?", userID).Where("user_roles.role_id IN (?)", roles.IDs()).Count(&count).Error
	return count > 0, err
}

// HasDirectPermission does the user have the given permission? (not including the permissios of the roles)
func (u *UserRepo) HasDirectPermission(userID string, permission models.Permission) (b bool, err error) {
	var count int64
	err = u.db.GetInstance().Table("user_permissions").Where("user_permissions.user_id = ?", userID).Where("user_permissions.permission_id = ?", permission.ID).Count(&count).Error
	return count > 0, err
}

func (u *UserRepo) HasAllDirectPermissions(userID string, permissions schema.Permission) (b bool, err error) {
	var count int64
	err = u.db.GetInstance().Table("user_permissions").Where("user_permissions.user_id = ?", userID).Where("user_permissions.permission_id IN (?)", permissions.IDs()).Count(&count).Error
	return permissions.Len() == count, err
}

func (u *UserRepo) HasAnyDirectPermissions(userID string, permissions schema.Permission) (b bool, err error) {
	var count int64
	err = u.db.GetInstance().Table("user_permissions").Where("user_permissions.user_id = ?", userID).Where("user_permissions.permission_id IN (?)", permissions.IDs()).Count(&count).Error
	return count > 0, err
}
