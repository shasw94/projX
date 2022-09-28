package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shasw94/projX/app/interfaces"
	"github.com/shasw94/projX/app/schema"
	"github.com/shasw94/projX/pkg/errors"
	gohttp "github.com/shasw94/projX/pkg/http/wrapper"
	"github.com/shasw94/projX/pkg/utils"
	"github.com/shasw94/projX/validation"
)

type RoleAPI struct {
	service interfaces.IRoleService
}

// NewRoleAPI return new RoleAPI pointer
func NewRoleAPI(service interfaces.IRoleService) *RoleAPI {
	return &RoleAPI{service: service}
}

// CreateRole create new role
func (r *RoleAPI) CreateRole(c *gin.Context) gohttp.Response {
	var params schema.RoleBodyParams
	if err := c.ShouldBindJSON(&params); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	validator := validation.New()
	if err := validator.ValidateStruct(params); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	ctx := c.Request.Context()
	user, err := r.service.Create(ctx, &params)
	if err != nil {
		return gohttp.Response{
			Error: err,
		}
	}

	var res schema.Role
	err = utils.Copy(&res, &user)
	if err != nil {
		return gohttp.Response{
			Error: err,
		}
	}
	return gohttp.Response{
		Data: res,
	}
}
