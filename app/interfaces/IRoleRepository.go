package interfaces

import (
	"github.com/shasw94/projX/app/models"
	"github.com/shasw94/projX/app/repositories/scopes"
	"github.com/shasw94/projX/app/schema"
)

// IRoleRepository interface
type IRoleRepository interface {
	GetByName(string) (*models.Role, error)
	Create(*models.Role) error
	GetRoleByID(string) (*models.Role, error)
	GetRoleByIDWithPermissions(string) (*models.Role, error)

	GetRoleByGuardName(string) (*models.Role, error)
	GetRoleByGuardNameWithPermissions(string) (*models.Role, error)

	// Multiple fetch options

	GetRoles([]string) (*schema.Roles, error)
	GetRolesWithPermissions([]string) (*schema.Roles, error)

	GetRolesByGuardNames([]string) (*schema.Roles, error)
	GetRolesByGuardNamesWithPermissions([]string) (*schema.Roles, error)

	// ID fetch options

	GetRoleIDs(scopes.GormPager) ([]string, int64, error)
	GetRoleIDsOfUser(string, scopes.GormPager) ([]string, int64, error)
	GetRoleIDsOfPermission(string, scopes.GormPager) ([]string, int64, error)

	// FirstOrCreate & Updates & Delete

	FirstOrCreate(*models.Role) error
	Updates(*models.Role, map[string]interface{}) error
	Delete(*models.Role) error

	// Actions

	AddPermissions(*models.Role, *schema.Permission) error
	ReplacePermissions(*models.Role, *schema.Permission) error
	RemovePermissions(*models.Role, *schema.Permission) error
	ClearPermissions(*models.Role) error

	// Controls

	HasPermission(*schema.Roles, *models.Permission) (bool, error)
	HasAllPermissions(*schema.Roles, *schema.Permission) (bool, error)
	HasAnyPermissions(*schema.Roles, *schema.Permission) (bool, error)
}
