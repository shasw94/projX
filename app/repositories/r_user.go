package repositories

import (
	"github.com/jinzhu/copier"
	"github.com/shasw94/projX/app/interfaces"
	"github.com/shasw94/projX/app/models"
	"github.com/shasw94/projX/app/schema"
	"github.com/shasw94/projX/pkg/errors"
	"github.com/shasw94/projX/pkg/utils"
	"golang.org/x/crypto/bcrypt"
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
