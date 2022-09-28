package interfaces

import (
	"context"
	"github.com/shasw94/projX/app/models"
	"github.com/shasw94/projX/app/schema"
)

type IRoleService interface {
	Create(ctx context.Context, item *schema.RoleBodyParams) (*models.Role, error)
}
