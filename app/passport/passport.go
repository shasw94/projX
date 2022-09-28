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
		err = p.roleRepo.AddPermissions(role, permissions)
	}

	return
}
