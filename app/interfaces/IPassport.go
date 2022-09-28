package interfaces

import (
	"github.com/shasw94/projX/app/models"
	"github.com/shasw94/projX/app/schema"
)

type IPassport interface {
	GetRole(r interface{}, withPermissions bool) (*models.Role, error)
	GetRoles(r interface{}, withPermissions bool) (*schema.Roles, error)
	GetAllRoles(option *schema.RoleOption) (roles *schema.Roles, totalCount int64, err error)
	GetRolesOfUser(userID string, option *schema.RoleOption) (roles *schema.Roles, totalCount int64, err error)
	CreateRole(name string, description string) error
	DeleteRole(r interface{}) error
	AddPermissionsToRole(r interface{}, p interface{}) error
	ReplacePermissionsToRole(r interface{}, p interface{}) error
	RemovePermissionsFromRole(r interface{}, p interface{}) error
	GetPermission(p interface{}) (permission *models.Permission, err error)
	GetPermissions(p interface{}) (permissions *schema.Permission, err error)
	GetAllPermissions(option *schema.PermissionOption) (permissions *schema.Permission, totalCount int64, err error)
	GetDirectPermissionsOfUser(userID uint, option *schema.PermissionOption) (permissions *schema.Permission, totalCount int64, err error)
	GetPermissionsOfRoles(r interface{}, option *schema.PermissionOption) (permissions *schema.Permission, totalCount int64, err error)
	GetAllPermissionsOfUser(userID string) (permissions *schema.Permission, err error)
	CreatePermission(name string, description string) error
	DeletePermission(p interface{}) error
	AddPermissionsToUser(userID string, p interface{}) error
	ReplacePermissionsToUser(userID string, p interface{}) error
	RemovePermissionsFromUser(userID string, p interface{}) error
	AddRolesToUser(userID string, r interface{}) error
	ReplaceRolesToUser(userID string, r interface{}) error
	RemoveRolesFromUser(userID string, r interface{}) error
	RoleHasPermission(r interface{}, p interface{}) (b bool, err error)
	RoleHasAllPermissions(r interface{}, p interface{}) (b bool, err error)
	RoleHasAnyPermissions(r interface{}, p interface{}) (b bool, err error)
	UserHasRole(userID string, r interface{}) (b bool, err error)
	UserHasAllRoles(userID string, r interface{}) (b bool, err error)
	UserHasAnyRoles(userID string, r interface{}) (b bool, err error)
	UserHasDirectPermission(userID string, p interface{}) (b bool, err error)
	UserHasAllDirectPermissions(userID string, p interface{}) (b bool, err error)
	UserHasAnyDirectPermissions(userID string, p interface{}) (b bool, err error)
	UserHasPermission(userID string, p interface{}) (b bool, err error)
	UserHasAllPermissions(userID string, p interface{}) (b bool, err error)
}
