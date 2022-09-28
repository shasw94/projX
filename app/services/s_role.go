package services

import (
	"context"
	"github.com/shasw94/projX/app/interfaces"
	"github.com/shasw94/projX/app/models"
	"github.com/shasw94/projX/app/schema"
	"github.com/shasw94/projX/pkg/utils"
)

// RoleService role service
type RoleService struct {
	repo interfaces.IRoleRepository
}

// NewRoleService return new IRoleService interface
func NewRoleService(repo interfaces.IRoleRepository) interfaces.IRoleService {
	return &RoleService{repo: repo}
}

// Create creates new role
func (r *RoleService) Create(ctx context.Context, item *schema.RoleBodyParams) (*models.Role, error) {
	var role models.Role
	err := utils.Copy(&role, &item)
	if err != nil {
		return nil, err
	}

	err = r.repo.Create(&role)
	if err != nil {
		return nil, err
	}

	return &role, nil
}