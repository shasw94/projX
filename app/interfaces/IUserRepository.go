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
}
