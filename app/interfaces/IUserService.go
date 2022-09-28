package interfaces

import (
	"context"
	"github.com/shasw94/projX/app/models"
	"github.com/shasw94/projX/app/schema"
)

type IUserService interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	List(ctx context.Context, param *schema.UserQueryParam) (*[]models.User, error)
}
