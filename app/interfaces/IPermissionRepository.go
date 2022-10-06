package interfaces

import (
	"github.com/shasw94/projX/app/models"
	"github.com/shasw94/projX/app/repositories/scopes"
	"github.com/shasw94/projX/app/schema"
)

type IPermissionRepository interface {
	Migratable

	GetPermissionByID(string) (models.Permission, error)
	GetPermissionByGuardName(string) (models.Permission, error)

	// Multiple fetch operations
	GetPermissions([]string) (schema.Permission, error)
	GetPermissionsByGuardNames([]string) (schema.Permission, error)

	// ID fetch options
	GetPermissionIDs(scopes.GormPager) ([]string, int64, error)
	GetDirectPermissionIDsOfUserByID(string, scopes.GormPager) ([]string, int64, error)
	GetPermissionIDsOfRolesByIDs([]string, scopes.GormPager) ([]string, int64, error)

	// FirstOrCreate & Updates & Delete
	FirstOrCreate(permission *models.Permission) (err error)
	Updates(permission *models.Permission, updates map[string]interface{}) (err error)
	Delete(permission *models.Permission) (err error)
}

type Migratable interface {
	Migrate() error
}
