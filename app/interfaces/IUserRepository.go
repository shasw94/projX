package interfaces

import (
	"github.com/shasw94/projX/app/models"
	"github.com/shasw94/projX/app/schema"
)

type IUserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetUserByToken(token string) (*models.User, error)
	List(queryParam *schema.UserQueryParam) (*[]models.User, error)
	Login(item *schema.LoginBodyParams) (*models.User, error)
	RemoveToken(userID string) (*models.User, error)
	Update(userID string, bodyParam *schema.UserUpdateBodyParam) (*models.User, error)
	AddPermissions(userID string, permissions schema.Permission) (err error)
	ReplacePermissions(userID string, permissions schema.Permission) (err error)
	RemovePermissions(userID string, permissiosn schema.Permission) (err error)
	AddRoles(userID string, roles schema.Roles) (err error)
	ReplaceRoles(userID string, roles schema.Roles) (err error)
	RemoveRoles(userID string, roles schema.Roles) (err error)
	ClearRoles(userID string) (err error)
	ClearPermissions(userID string) (err error)
	HasRole(userID string, role models.Role) (b bool, err error)
	HasAllRoles(userID string, roles schema.Roles) (b bool, err error)
	HasAnyRoles(userID string, roles schema.Roles) (b bool, err error)
	HasDirectPermission(userID string, permission models.Permission) (b bool, err error)
	HasAllDirectPermissions(userID string, permissions schema.Permission) (b bool, err error)
	HasAnyDirectPermissions(userID string, permissions schema.Permission) (b bool, err error)
}
