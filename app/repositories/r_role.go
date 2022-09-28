package repositories

import (
	"github.com/shasw94/projX/app/interfaces"
	"github.com/shasw94/projX/app/models"
	"github.com/shasw94/projX/app/models/pivot"
	"github.com/shasw94/projX/app/repositories/scopes"
	"github.com/shasw94/projX/app/schema"
	"github.com/shasw94/projX/pkg/errors"
	"gorm.io/gorm"
)

type RoleRepo struct {
	db interfaces.IDatabase
}

// NewRoleRepository return new IRoleRepository interface
func NewRoleRepository(db interfaces.IDatabase) interfaces.IRoleRepository {
	return &RoleRepo{db: db}
}

// Create new role
func (r *RoleRepo) Create(role *models.Role) error {
	if err := r.db.GetInstance().Model(&models.Role{}).Create(&role).Error; err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}
	return nil
}

// GetByName get role by name
func (r *RoleRepo) GetByName(name string) (*models.Role, error) {
	var role models.Role
	if err := r.db.GetInstance().Model(&models.Role{}).Where("name = ?", name).First(&role).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return &role, nil
}

func (r *RoleRepo) Migrate() error {
	err := r.db.GetInstance().AutoMigrate(models.Role{})
	err = r.db.GetInstance().AutoMigrate(pivot.UserRole{})
	return err
}

func (r *RoleRepo) GetRoleByID(ID string) (*models.Role, error) {
	role := models.Role{}
	err := r.db.GetInstance().First(&role, "roles.id = ?", ID).Error
	if err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return &role, err
}

// GetRoleByIDWithPermissions get role by id with its permissions.
// @param uint
// @return models.Role, error
func (r *RoleRepo) GetRoleByIDWithPermissions(ID string) (*models.Role, error) {
	role := models.Role{}
	err := r.db.GetInstance().Preload("Permissions").First(&role, "roles.id = ?", ID).Error
	if err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return &role, err
}

// GetRoleByGuardName get role by guard name.
// @param string
// @return models.Role, error
func (r *RoleRepo) GetRoleByGuardName(guardName string) (*models.Role, error) {
	role := models.Role{}
	err := r.db.GetInstance().Where("roles.guard_name = ?", guardName).First(&role).Error
	if err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return &role, err
}

// GetRoleByGuardNameWithPermissions get role by guard name with its permissions.
// @param string
// @return models.Role, error
func (r *RoleRepo) GetRoleByGuardNameWithPermissions(guardName string) (*models.Role, error) {
	role := models.Role{}
	err := r.db.GetInstance().Preload("Permissions").Where("roles.guard_name = ?", guardName).First(&role).Error
	if err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return &role, err
}

// MULTIPLE FETCH OPTIONS

// GetRoles get roles by ids.
// @param []uint
// @return *schema.Roles, error
func (r *RoleRepo) GetRoles(IDs []string) (*schema.Roles, error) {
	roles := schema.Roles{}
	err := r.db.GetInstance().Where("roles.id IN (?)", IDs).Find(&roles).Error
	if err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return &roles, err
}

// GetRolesWithPermissions get roles by ids with its permissions.
// @param []uint
// @return *schema.Roles, error
func (r *RoleRepo) GetRolesWithPermissions(IDs []string) (*schema.Roles, error) {
	roles := schema.Roles{}
	err := r.db.GetInstance().Preload("Permissions").Where("roles.id IN (?)", IDs).Find(&roles).Error
	if err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return &roles, err
}

// GetRolesByGuardNames get roles by guard names.
// @param []string
// @return *schema.Roles, error
func (r *RoleRepo) GetRolesByGuardNames(guardNames []string) (*schema.Roles, error) {
	roles := schema.Roles{}
	err := r.db.GetInstance().Where("roles.guard_name IN (?)", guardNames).Find(&roles).Error
	if err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return &roles, err
}

// GetRolesByGuardNamesWithPermissions get roles by guard names.
// @param []string
// @return *schema.Roles, error
func (r *RoleRepo) GetRolesByGuardNamesWithPermissions(guardNames []string) (*schema.Roles, error) {
	roles := schema.Roles{}
	err := r.db.GetInstance().Preload("Permissions").Where("roles.guard_name IN (?)", guardNames).Find(&roles).Error
	if err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return &roles, err
}

// ID FETCH OPTIONS

// GetRoleIDs get role ids. (with pagination)
// @param repositories_scopes.GormPager
// @return []uint, int64, error
func (r *RoleRepo) GetRoleIDs(pagination scopes.GormPager) (roleIDs []string, totalCount int64, err error) {
	err = r.db.GetInstance().Model(&models.Role{}).Count(&totalCount).Scopes(r.paginate(pagination)).Pluck("roles.id", &roleIDs).Error
	if err != nil {
		return nil, 0, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return
}

// GetRoleIDsOfUser get role ids of user. (with pagination)
// @param uint
// @param repositories_scopes.GormPager
// @return []uint, int64, error
func (r *RoleRepo) GetRoleIDsOfUser(userID string, pagination scopes.GormPager) (roleIDs []string, totalCount int64, err error) {
	err = r.db.GetInstance().Table("user_roles").Where("user_roles.user_id = ?", userID).Count(&totalCount).Scopes(r.paginate(pagination)).Pluck("user_roles.role_id", &roleIDs).Error
	if err != nil {
		return nil, 0, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return
}

// GetRoleIDsOfPermission get role ids of permission. (with pagination)
// @param uint
// @param repositories_scopes.GormPager
// @return []uint, int64, error
func (r *RoleRepo) GetRoleIDsOfPermission(permissionID string, pagination scopes.GormPager) (roleIDs []string, totalCount int64, err error) {
	err = r.db.GetInstance().Table("role_permissions").Where("role_permissions.permission_id = ?", permissionID).Count(&totalCount).Scopes(r.paginate(pagination)).Pluck("role_permissions.role_id", &roleIDs).Error
	if err != nil {
		return nil, 0, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	return
}

// FirstOrCreate & Updates & Delete

// FirstOrCreate create new role if name not exist.
// @param *models.Role
// @return error
func (r *RoleRepo) FirstOrCreate(role *models.Role) error {
	err := r.db.GetInstance().Where(models.Role{GuardName: role.GuardName}).FirstOrCreate(role).Error
	if err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}
	return nil
}

// Updates update role.
// @param *models.Role
// @param map[string]interface{}
// @return error
func (r *RoleRepo) Updates(role *models.Role, updates map[string]interface{}) error {
	err := r.db.GetInstance().Model(role).Updates(updates).Error
	if err != nil {
		return errors.ErrorDatabaseUpdate.Newm(err.Error())
	}
	return nil
}

// Delete delete role.
// @param *models.Role
// @return error
func (r *RoleRepo) Delete(role *models.Role) (err error) {
	return r.db.GetInstance().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_roles.role_id = ?", role.ID).Delete(&pivot.UserRole{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Delete(role).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

// ACTIONS

// AddPermissions add permissions to role.
// @param *models.Role
// @param *schema.Permission
// @return error
func (r *RoleRepo) AddPermissions(role *models.Role, permissions *schema.Permission) error {
	return r.db.GetInstance().Model(role).Association("Permissions").Append(permissions.Origin())
}

// ReplacePermissions replace permissions of role.
// @param *models.Role
// @param *schema.Permission
// @return error
func (r *RoleRepo) ReplacePermissions(role *models.Role, permissions *schema.Permission) error {
	return r.db.GetInstance().Model(role).Association("Permissions").Replace(permissions.Origin())
}

// RemovePermissions remove permissions of role.
// @param *models.Role
// @param *schema.Permission
// @return error
func (r *RoleRepo) RemovePermissions(role *models.Role, permissions *schema.Permission) error {
	return r.db.GetInstance().Model(role).Association("Permissions").Delete(permissions.Origin())
}

// ClearPermissions remove all permissions of role.
// @param *models.Role
// @return error
func (r *RoleRepo) ClearPermissions(role *models.Role) (err error) {
	return r.db.GetInstance().Model(role).Association("Permissions").Clear()
}

// Controls

// HasPermission does the role or any of the roles have given permission?
// @param *schema.Roles
// @param models.Permission
// @return bool, error
func (r *RoleRepo) HasPermission(roles *schema.Roles, permission *models.Permission) (b bool, err error) {
	var count int64
	err = r.db.GetInstance().Table("role_permissions").Where("role_permissions.role_id IN (?)", roles.IDs()).Where("role_permissions.permission_id = ?", permission.ID).Count(&count).Error
	return count > 0, err
}

// HasAllPermissions does the role or roles have all the given permissions?
// @param *schema.Roles
// @param *schema.Permission
// @return bool, error
func (r *RoleRepo) HasAllPermissions(roles *schema.Roles, permissions *schema.Permission) (b bool, err error) {
	var count int64
	err = r.db.GetInstance().Table("role_permissions").Where("role_permissions.role_id IN (?)", roles.IDs()).Where("role_permissions.permission_id IN (?)", permissions.IDs()).Count(&count).Error
	return roles.Len()*permissions.Len() == count, err
}

// HasAnyPermissions does the role or roles have any of the given permissions?
// @param *schema.Roles
// @param *schema.Permission
// @return bool, error
func (r *RoleRepo) HasAnyPermissions(roles *schema.Roles, permissions *schema.Permission) (b bool, err error) {
	var count int64
	err = r.db.GetInstance().Table("role_permissions").Where("role_permissions.role_id IN (?)", roles.IDs()).Where("role_permissions.permission_id IN (?)", permissions.IDs()).Count(&count).Error
	return count > 0, err
}

// paginate paging if pagination option is true.
// @param repositories_scopes.GormPager
// @return func(db *gorm.DB) *gorm.DB
func (r *RoleRepo) paginate(pagination scopes.GormPager) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pagination != nil {
			db.Scopes(pagination.ToPaginate())
		}

		return db
	}
}
