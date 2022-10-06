package passport

import (
	"github.com/shasw94/projX/app/interfaces"
	"github.com/shasw94/projX/app/models"
	"github.com/shasw94/projX/app/repositories/scopes"
	"github.com/shasw94/projX/app/schema"
	"github.com/shasw94/projX/pkg/errors"
	"github.com/shasw94/projX/pkg/utils"
)

var errUnsupportedValueType = errors.New("err unsupported value type")

type Passport struct {
	roleRepo interfaces.IRoleRepository
	userRepo interfaces.IUserRepository
	permRepo interfaces.IPermissionRepository
}

func NewPassport(roleRepo interfaces.IRoleRepository, userRepo interfaces.IUserRepository) *Passport {
	return &Passport{roleRepo: roleRepo, userRepo: userRepo}
}

func (p *Passport) GetRole(r interface{}, withPermissions bool) (*models.Role, error) {
	if utils.IsArray(r) {
		//var roles []models.Role
		roles, err := p.GetRoles(r, withPermissions)
		if err != nil {
			return nil, errors.ErrorDatabaseGet.Newm(err.Error())
		}
		if len(*roles) > 0 {
			p := *roles
			role := p[0]
			return &role, nil
		}
		return nil, errors.ErrorDatabaseGet.Newm("No roles found")
	}

	if utils.IsString(r) {
		if withPermissions {
			return p.roleRepo.GetRoleByGuardNameWithPermissions(utils.Guard(r.(string)))
		}
		return p.roleRepo.GetRoleByGuardName(utils.Guard(r.(string)))
	}

	return nil, errors.ErrorExistRole.Newm("unsupported value type")
}

func (p *Passport) GetRoles(r interface{}, withPermissions bool) (*schema.Roles, error) {
	if !utils.IsArray(r) {
		role, err := p.GetRole(r, withPermissions)
		if err != nil {
			return nil, err
		}
		roles := schema.Roles{*role}
		return &roles, nil
	}

	if utils.IsStringArray(r) {
		if withPermissions {
			return p.roleRepo.GetRolesWithPermissions(r.([]string))
		}
		return p.roleRepo.GetRoles(r.([]string))
	}
	return nil, errors.ErrorExistRole.Newm("unsupported value type")
}

func (p *Passport) GetAllRoles(option schema.RoleOption) (roles *schema.Roles, totalCount int64, err error) {
	var roleIDs []string
	if option.Pagination == nil {
		roleIDs, totalCount, err = p.roleRepo.GetRoleIDs(nil)
	} else {
		roleIDs, totalCount, err = p.roleRepo.GetRoleIDs(&scopes.GormPagination{Pagination: option.Pagination.Get()})
	}
	roles, err = p.GetRoles(roleIDs, option.WithPermissions)
	return
}

func (p *Passport) GetRolesOfUser(userID string, option schema.RoleOption) (*schema.Roles, int64, error) {
	var roleIDs []string
	var totalCount int64
	var err error
	if option.Pagination == nil {
		roleIDs, totalCount, err = p.roleRepo.GetRoleIDsOfUser(userID, nil)
	} else {
		roleIDs, totalCount, err = p.roleRepo.GetRoleIDsOfUser(userID, &scopes.GormPagination{Pagination: option.Pagination.Get()})
	}

	roles, err := p.GetRoles(roleIDs, option.WithPermissions)
	return roles, totalCount, err
}

func (p *Passport) CreateRole(name, description string) (err error) {
	return p.roleRepo.FirstOrCreate(&models.Role{
		Name:        name,
		GuardName:   utils.Guard(name),
		Description: description,
	})
}

// DeleteRole delete role.
// If the role is in use, its relations from the pivot tables are deleted.
// First parameter can be role name or id.
// @param interface{}
// @return error
func (p *Passport) DeleteRole(r interface{}) (err error) {
	role, err := p.GetRole(r, false)
	if err != nil {
		return err
	}
	return p.roleRepo.Delete(role)
}

// AddPermissionsToRole add permission to role.
// First parameter can be role name or id, second parameter can be permission name(s) or id(s).
// If the first parameter is an array, the first element of the first parameter is used.
// @param interface{}
// @param interface{}
// @return error
func (p *Passport) AddPermissionsToRole(r interface{}, per interface{}) (err error) {
	role, err := p.GetRole(r, false)
	if err != nil {
		return err
	}

	permissions, err := p.GetPermissions(per)
	if err != nil {
		return err
	}

	if permissions.Len() > 0 {
		err = p.roleRepo.AddPermissions(role, &permissions)
	}

	return
}

func (p *Passport) ReplacePermissionsToRole(r interface{}, s interface{}) (err error) {
	role, err := p.GetRole(r, false)
	if err != nil {
		return err
	}
	permissions, err := p.GetPermissions(s)
	if err != nil {
		return err
	}
	if permissions.Len() > 0 {
		return p.roleRepo.ReplacePermissions(role, &permissions)
	}

	return p.roleRepo.ClearPermissions(role)
}

func (p *Passport) RemovePermissionsFromRole(r interface{}, s interface{}) (err error) {
	role, err := p.GetRole(r, false)
	if err != nil {
		return err
	}

	permissions, err := p.GetPermissions(s)
	if err != nil {
		return err
	}
	if permissions.Len() > 0 {
		err = p.roleRepo.RemovePermissions(role, &permissions)
	}

	return
}

// PERMISSION

func (p *Passport) GetPermission(s interface{}) (permission models.Permission, err error) {
	if utils.IsArray(s) {
		var permissions []models.Permission
		permissions, err = p.GetPermissions(p)
		if err != nil {
			return models.Permission{}, err
		}
		if len(permissions) > 0 {
			permission = permissions[0]
		}
		return
	}

	if utils.IsString(s) {
		return p.permRepo.GetPermissionByID(utils.Guard(s.(string)))
	}

	return models.Permission{}, errors.ErrorExistRole.Newm("unsupported value type")
}

// GetPermissions fetch permissions according to the permission names or ids.
// First parameter is can be permission name(s) or id(s).
// @param interface{}
// @return collections.Permission, error
func (s *Passport) GetPermissions(p interface{}) (permissions schema.Permission, err error) {
	if !utils.IsArray(p) {
		var permission models.Permission
		permission, err = s.GetPermission(p)
		if err != nil {
			return schema.Permission{}, err
		}
		permissions = schema.Permission{permission}
		return
	}

	if utils.IsStringArray(p) {
		return s.permRepo.GetPermissions(p.([]string))
	}

	return schema.Permission{}, errUnsupportedValueType
}

// GetAllPermissions fetch all the permissions. (with pagination option).
// First parameter is permission option.
// @param options.PermissionOption
// @return collections.Permission, int64, error
func (s *Passport) GetAllPermissions(option schema.PermissionOption) (permissions schema.Permission, totalCount int64, err error) {
	var permissionIDs []string
	if option.Pagination == nil {
		permissionIDs, totalCount, err = s.permRepo.GetPermissionIDs(nil)
	} else {
		permissionIDs, totalCount, err = s.permRepo.GetPermissionIDs(&scopes.GormPagination{option.Pagination.Get()})
	}
	permissions, err = s.GetPermissions(permissionIDs)
	return
}

// GetDirectPermissionsOfUser fetch all direct permissions of the user. (with pagination option)
// First parameter is user id, second parameter is permission option.
// @param uint
// @param options.PermissionOption
// @return collections.Permission, int64, error
func (s *Passport) GetDirectPermissionsOfUser(userID string, option schema.PermissionOption) (permissions schema.Permission, totalCount int64, err error) {
	var permissionIDs []string
	if option.Pagination == nil {
		permissionIDs, totalCount, err = s.permRepo.GetDirectPermissionIDsOfUserByID(userID, nil)
	} else {
		permissionIDs, totalCount, err = s.permRepo.GetDirectPermissionIDsOfUserByID(userID, &scopes.GormPagination{option.Pagination.Get()})
	}
	permissions, err = s.GetPermissions(permissionIDs)
	return
}

// GetPermissionsOfRoles fetch all permissions of the roles. (with pagination option)
// First parameter can be role name(s) or id(s), second parameter is permission option.
// @param interface{}
// @param options.PermissionOption
// @return collections.Permission, int64, error
func (s *Passport) GetPermissionsOfRoles(r interface{}, option schema.PermissionOption) (permissions schema.Permission, totalCount int64, err error) {
	var roles *schema.Roles
	roles, err = s.GetRoles(r, false)
	if err != nil {
		return schema.Permission{}, 0, err
	}

	var permissionIDs []string
	if option.Pagination == nil {
		permissionIDs, totalCount, err = s.permRepo.GetPermissionIDsOfRolesByIDs(roles.IDs(), nil)
	} else {
		permissionIDs, totalCount, err = s.permRepo.GetPermissionIDsOfRolesByIDs(roles.IDs(), &scopes.GormPagination{option.Pagination.Get()})
	}
	permissions, err = s.GetPermissions(permissionIDs)
	return
}

// GetAllPermissinosOfUser fetch all permissions of the user that come with direct and roles
// @param string
// @return schema.Permission, error
func (s *Passport) GetAllPermissionsOfUser(userID string) (permissions schema.Permission, err error) {
	var userRoleIDs []string
	userRoleIDs, _, err = s.roleRepo.GetRoleIDsOfUser(userID, nil)
	if err != nil {
		return schema.Permission{}, err
	}

	var rolePermissionIDs []string
	rolePermissionIDs, _, err = s.permRepo.GetPermissionIDsOfRolesByIDs(userRoleIDs, nil)
	if err != nil {
		return schema.Permission{}, err
	}

	var userDirectPermissionIDs []string
	userDirectPermissionIDs, _, err = s.permRepo.GetDirectPermissionIDsOfUserByID(userID, nil)
	if err != nil {
		return schema.Permission{}, err
	}
	return s.GetPermissions(utils.RemoveDuplicateValues(utils.JoinStringArrays(rolePermissionIDs, userDirectPermissionIDs)))
}

// CreatePermission create new permission
// Name parameter is converted to guard name. example: create $#% contact -> create-contact
// If a permission with the same name has been created before, it will not create it again
// @param string
// @param string
// @return error
func (s *Passport) CreatePermission(name string, description string) (err error) {
	return s.permRepo.FirstOrCreate(&models.Permission{
		Name:        name,
		GuardName:   utils.Guard(name),
		Description: description,
	})
}

// DeletePermission delete permission
// If the permission is in use, its relations from the pivot tables are deleted
// First parameter can be permission name or id
// If the first paramter is an array, the first element of the given array is used.
// @param interface{}
// @return error
func (s *Passport) DeletePermission(p interface{}) (err error) {
	var permission models.Permission
	permission, err = s.GetPermission(p)
	if err != nil {
		return err
	}
	return s.permRepo.Delete(&permission)
}

// ------------------ USER -------------------

// AddPermissionToUser add direct permission or permissions to user according to the permission names or ids.
// FIrst parameter is the user id, secodn parameter can be permission name(s) or id(s)
// @param uint
// @param interface{}
// @return error
func (s *Passport) AddPermissionToUser(userID string, p interface{}) (err error) {
	var permissions schema.Permission
	permissions, err = s.GetPermissions(p)
	if err != nil {
		return err
	}
	if permissions.Len() > 0 {
		err = s.userRepo.AddPermissions(userID, permissions)
	}

	return
}

func (p *Passport) ReplacePermissionsToUser(userID string, s interface{}) (err error) {
	permissions, err := p.GetPermissions(s)
	if err != nil {
		return err
	}

	if permissions.Len() > 0 {
		return p.userRepo.ReplacePermissions(userID, permissions)
	}
	return p.userRepo.ClearPermissions(userID)
}

func (p *Passport) RemovePermissionsFromUser(userId string, s interface{}) (err error) {
	permissions, err := p.GetPermissions(s)
	if err != nil {
		return err
	}

	if permissions.Len() > 0 {
		err = p.userRepo.RemovePermissions(userId, permissions)
	}

	return
}

func (p *Passport) AddRolesToUser(userID string, r interface{}) (err error) {
	roles, err := p.GetRoles(r, false)
	if err != nil {
		return err
	}

	if roles.Len() > 0 {
		err = p.userRepo.AddRoles(userID, *roles)
	}

	return
}

func (p *Passport) ReplaceRolesToUser(userID string, r interface{}) (err error) {
	roles, err := p.GetRoles(r, false)
	if err != nil {
		return err
	}

	if roles.Len() > 0 {
		return p.userRepo.ReplaceRoles(userID, *roles)
	}

	return p.userRepo.ClearRoles(userID)
}

func (p *Passport) RemoveRolesFromUser(userID string, r interface{}) (err error) {
	roles, err := p.GetRoles(r, false)
	if err != nil {
		return err
	}

	if roles.Len() > 0 {
		err = p.userRepo.RemoveRoles(userID, *roles)
	}

	return
}

func (p *Passport) RoleHasPermission(r interface{}, s interface{}) (b bool, err error) {
	roles, err := p.GetRoles(r, false)
	if err != nil {
		return false, err
	}

	permission, err := p.GetPermission(s)
	if err != nil {
		return false, err
	}

	return p.roleRepo.HasPermission(roles, &permission)
}

//---------

// RoleHasAllPermissions does the role or roles have all the given permissions?
// First parameter is can be role name(s) or id(s), second parameter is can be permission name(s) or id(s).
// @param interface{}
// @param interface{}
// @return error
func (p *Passport) RoleHasAllPermissions(r interface{}, s interface{}) (b bool, err error) {
	roles, err := p.GetRoles(r, false)
	if err != nil {
		return false, err
	}

	permissions, err := p.GetPermissions(p)
	if err != nil {
		return false, err
	}

	return p.roleRepo.HasAllPermissions(roles, &permissions)
}

// RoleHasAnyPermissions does the role or roles have any of the given permissions?
// First parameter is can be role name(s) or id(s), second parameter is can be permission name(s) or id(s).
// @param interface{}
// @param interface{}
// @return error
func (p *Passport) RoleHasAnyPermissions(r interface{}, s interface{}) (b bool, err error) {
	roles, err := p.GetRoles(r, false)
	if err != nil {
		return false, err
	}

	permissions, err := p.GetPermissions(s)
	if err != nil {
		return false, err
	}

	return p.roleRepo.HasAnyPermissions(roles, &permissions)
}

// USER

// UserHasRole does the user have the given role?
// First parameter is the user id, second parameter is can be role name or id.
// If the second parameter is an array, the first element of the given array is used.
// @param uint
// @param interface{}
// @return bool, error
func (p *Passport) UserHasRole(userID string, r interface{}) (b bool, err error) {
	role, err := p.GetRole(r, false)
	if err != nil {
		return false, err
	}
	return p.userRepo.HasRole(userID, *role)
}

// UserHasAllRoles does the user have all the given roles?
// First parameter is the user id, second parameter is can be role name(s) or id(s).
// @param uint
// @param interface{}
// @return bool, error
func (p *Passport) UserHasAllRoles(userID string, r interface{}) (b bool, err error) {
	roles, err := p.GetRoles(r, false)
	if err != nil {
		return false, err
	}
	return p.userRepo.HasAllRoles(userID, *roles)
}

// UserHasAnyRoles does the user have any of the given roles?
// First parameter is the user id, second parameter is can be role name(s) or id(s).
// @param uint
// @param interface{}
// @return bool, error
func (p *Passport) UserHasAnyRoles(userID string, r interface{}) (b bool, err error) {
	roles, err := p.GetRoles(r, false)
	if err != nil {
		return false, err
	}
	return p.userRepo.HasAnyRoles(userID, *roles)
}

// UserHasDirectPermission does the user have the given permission? (not including the permissions of the roles)
// First parameter is the user id, second parameter is can be permission name or id.
// If the second parameter is an array, the first element of the given array is used.
// @param uint
// @param interface{}
// @return bool, error
func (p *Passport) UserHasDirectPermission(userID string, s interface{}) (b bool, err error) {
	permission, err := p.GetPermission(s)
	if err != nil {
		return false, err
	}
	return p.userRepo.HasDirectPermission(userID, permission)
}

// UserHasAllDirectPermissions does the user have all the given permissions? (not including the permissions of the roles)
// First parameter is the user id, second parameter is can be permission name(s) or id(s).
// @param uint
// @param interface{}
// @return bool, error
func (p *Passport) UserHasAllDirectPermissions(userID string, s interface{}) (b bool, err error) {
	permissions, err := p.GetPermissions(s)
	if err != nil {
		return false, err
	}
	return p.userRepo.HasAllDirectPermissions(userID, permissions)
}

// UserHasAnyDirectPermissions does the user have any of the given permissions? (not including the permissions of the roles)
// First parameter is the user id, second parameter is can be permission name(s) or id(s).
// @param uint
// @param interface{}
// @return bool, error
func (p *Passport) UserHasAnyDirectPermissions(userID string, s interface{}) (b bool, err error) {
	permissions, err := p.GetPermissions(s)
	if err != nil {
		return false, err
	}
	return p.userRepo.HasAnyDirectPermissions(userID, permissions)
}

// UserHasPermission does the user have the given permission? (including the permissions of the roles)
// First parameter is the user id, second parameter is can be permission name or id.
// If the second parameter is an array, the first element of the given array is used.
// @param uint
// @param interface{}
// @return bool, error
func (p *Passport) UserHasPermission(userID string, s interface{}) (b bool, err error) {
	permission, err := p.GetPermission(s)
	if err != nil {
		return false, err
	}

	var directPermissionIDs []string
	directPermissionIDs, _, err = p.permRepo.GetDirectPermissionIDsOfUserByID(userID, nil)
	if err != nil {
		return false, err
	}

	if utils.InArray(permission.ID, directPermissionIDs) {
		return true, err
	}

	var roleIDs []string
	roleIDs, _, err = p.roleRepo.GetRoleIDsOfUser(userID, nil)
	if err != nil {
		return false, err
	}

	var permissionIDs []string
	permissionIDs, _, err = p.permRepo.GetPermissionIDsOfRolesByIDs(roleIDs, nil)
	if err != nil {
		return false, err
	}

	if utils.InArray(permission.ID, permissionIDs) {
		return true, err
	}

	return false, err
}

// UserHasAllPermissions does the user have all the given permissions? (including the permissions of the roles).
// First parameter is the user id, second parameter is can be permission name(s) or id(s).
// @param uint
// @param interface{}
// @return bool, error
func (p *Passport) UserHasAllPermissions(userID string, s interface{}) (b bool, err error) {
	permissions, err := p.GetPermissions(s)
	if err != nil {
		return false, err
	}

	var userPermissionIDs []string
	userPermissionIDs, _, err = p.permRepo.GetDirectPermissionIDsOfUserByID(userID, nil)
	if err != nil {
		return false, err
	}

	var roleIDs []string
	roleIDs, _, err = p.roleRepo.GetRoleIDsOfUser(userID, nil)
	if err != nil {
		return false, err
	}

	var rolePermissionIDs []string
	rolePermissionIDs, _, err = p.permRepo.GetPermissionIDsOfRolesByIDs(roleIDs, nil)
	if err != nil {
		return false, err
	}

	allPermissionIDsOfUser := utils.RemoveDuplicateValues(utils.JoinStringArrays(userPermissionIDs, rolePermissionIDs))

	for _, permissionID := range permissions.IDs() {
		if !utils.InArray(permissionID, allPermissionIDsOfUser) {
			return false, err
		}
	}

	return true, err
}

// UserHasAnyPermissions does the user have any of the given permissions? (including the permissions of the roles).
// First parameter is the user id, second parameter is can be permission name(s) or id(s).
// @param uint
// @param interface{}
// @return bool, error
func (p *Passport) UserHasAnyPermissions(userID string, s interface{}) (b bool, err error) {
	permissions, err := p.GetPermissions(s)
	if err != nil {
		return false, err
	}

	var directPermissionIDs []string
	directPermissionIDs, _, err = p.permRepo.GetDirectPermissionIDsOfUserByID(userID, nil)
	if err != nil {
		return false, err
	}

	for _, permissionID := range permissions.IDs() {
		if utils.InArray(permissionID, directPermissionIDs) {
			return true, err
		}
	}

	var roleIDs []string
	roleIDs, _, err = p.roleRepo.GetRoleIDsOfUser(userID, nil)
	if err != nil {
		return false, err
	}

	var permissionIDs []string
	permissionIDs, _, err = p.permRepo.GetPermissionIDsOfRolesByIDs(roleIDs, nil)
	if err != nil {
		return false, err
	}

	for _, permissionID := range permissions.IDs() {
		if utils.InArray(permissionID, permissionIDs) {
			return true, err
		}
	}

	return false, err
}
