package repositories

import (
	"github.com/shasw94/projX/app/interfaces"
	"github.com/shasw94/projX/app/models"
	"github.com/shasw94/projX/app/models/pivot"
	"github.com/shasw94/projX/app/repositories/scopes"
	"github.com/shasw94/projX/app/schema"
	"gorm.io/gorm"
)

type PermissionRepo struct {
	db interfaces.IDatabase
}

func NewPermissionRepo(db interfaces.IDatabase) interfaces.IPermissionRepository {
	return &PermissionRepo{
		db: db,
	}
}

// Migrate generate tables from database
// @return error
func (repository *PermissionRepo) Migrate() (err error) {
	err = repository.db.GetInstance().AutoMigrate(models.Permission{})
	err = repository.db.GetInstance().AutoMigrate(pivot.UserPermission{})
	return
}

func (p *PermissionRepo) GetPermissionByID(ID string) (permission models.Permission, err error) {
	err = p.db.GetInstance().First(&permission, "permissions.id = ?", ID).Error
	return
}

// GetPermissionsByGuardName get permission by guard name.
// @param string
// @return models.Permission, error
func (repository *PermissionRepo) GetPermissionByGuardName(guardName string) (permission models.Permission, err error) {
	err = repository.db.GetInstance().Where("permissions.guard_name = ?", guardName).First(&permission).Error
	return
}

// MULTIPLE FETCH OPTIONS

// GetPermissions get permissions by ids.
// @param []string
// @return schema.Role, error
func (repository *PermissionRepo) GetPermissions(IDs []string) (permissions schema.Permission, err error) {
	err = repository.db.GetInstance().Where("permissions.id IN (?)", IDs).Find(&permissions).Error
	return
}

// GetPermissionsByGuardNames get permissions by guard names.
// @param []string
// @return schema.Permission, error
func (repository *PermissionRepo) GetPermissionsByGuardNames(guardNames []string) (permissions schema.Permission, err error) {
	err = repository.db.GetInstance().Where("permissions.guard_name IN (?)", guardNames).Find(&permissions).Error
	return
}

// ID FETCH OPTIONS

// GetPermissionIDs get permission ids. (with pagination)
// @param repositories_scopes.GormPager
// @return []string, int64, error
func (repository *PermissionRepo) GetPermissionIDs(pagination scopes.GormPager) (permissionIDs []string, totalCount int64, err error) {
	err = repository.db.GetInstance().Model(&models.Permission{}).Count(&totalCount).Scopes(repository.paginate(pagination)).Pluck("permissions.id", &permissionIDs).Error
	return
}

// GetDirectPermissionIDsOfUserByID get direct permission ids of user. (with pagination)
// @param string
// @param repositories_scopes.GormPager
// @return []string, int64, error
func (repository *PermissionRepo) GetDirectPermissionIDsOfUserByID(userID string, pagination scopes.GormPager) (permissionIDs []string, totalCount int64, err error) {
	err = repository.db.GetInstance().Table("user_permissions").Where("user_permissions.user_id = ?", userID).Count(&totalCount).Scopes(repository.paginate(pagination)).Pluck("user_permissions.permission_id", &permissionIDs).Error
	return
}

// GetPermissionIDsOfRolesByIDs get permission ids of roles. (with pagination)
// @param []string
// @param repositories_scopes.GormPager
// @return []string, int64, error
func (repository *PermissionRepo) GetPermissionIDsOfRolesByIDs(roleIDs []string, pagination scopes.GormPager) (permissionIDs []string, totalCount int64, err error) {
	err = repository.db.GetInstance().Table("role_permissions").Distinct("role_permissions.permission_id").Where("role_permissions.role_id IN (?)", roleIDs).Count(&totalCount).Scopes(repository.paginate(pagination)).Pluck("role_permissions.permission_id", &permissionIDs).Error
	return
}

// FirstOrCreate & Updates & Delete

// FirstOrCreate create new permission if name not exist.
// @param *models.Permission
// @return error
func (repository *PermissionRepo) FirstOrCreate(permission *models.Permission) error {
	return repository.db.GetInstance().Where(models.Role{GuardName: permission.GuardName}).FirstOrCreate(permission).Error
}

// Updates update permission.
// @param *models.Permission
// @param map[string]interface{}
// @return error
func (repository *PermissionRepo) Updates(permission *models.Permission, updates map[string]interface{}) (err error) {
	return repository.db.GetInstance().Model(permission).Updates(updates).Error
}

// Delete delete permission.
// @param *models.Permission
// @return error
func (repository *PermissionRepo) Delete(permission *models.Permission) (err error) {
	return repository.db.GetInstance().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_permissions.permission_id = ?", permission.ID).Delete(&pivot.UserPermission{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Delete(permission).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

// paginate pagging if pagination option is true.
// @param repositories_scopes.GormPager
// @return func(db *gorm.DB) *gorm.DB
func (repository *PermissionRepo) paginate(pagination scopes.GormPager) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pagination != nil {
			db.Scopes(pagination.ToPaginate())
		}

		return db
	}
}
