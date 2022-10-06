package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/shasw94/projX/app/interfaces"
	"github.com/shasw94/projX/app/schema"
	"github.com/shasw94/projX/logger"
	"github.com/shasw94/projX/pkg/errors"
	gohttp "github.com/shasw94/projX/pkg/http/wrapper"
	"github.com/shasw94/projX/pkg/utils"
)

// UserAPI handle user api
type UserAPI struct {
	service interfaces.IUserService
}

func NewUserAPI(service interfaces.IUserService) *UserAPI {
	return &UserAPI{service: service}
}

// GetByID get user by ID
func (u *UserAPI) GetByID(c *gin.Context) {
	userID := c.Param("id")
	ctx := c.Request.Context()
	// uid, err := strconv.ParseUint(userID, 10, 32)
	// if err != nil {
	// 	err = errors.Wrap(err, "API.GetByID")
	// 	logger.Error("Failed to get user: ", err)
	// 	c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), utils.ErrorGetDatabase))
	// 	return
	// }
	user, err := u.service.GetByID(ctx, userID)
	if err != nil {
		err = errors.Wrap(err, "API.GetByID")
		logger.Error("Failed to get user: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), utils.ErrorGetDatabase))
		return
	}

	var res schema.User
	copier.Copy(&res, &user)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

func (u *UserAPI) List(c *gin.Context) gohttp.Response {
	var queryParam schema.UserQueryParam
	if err := c.ShouldBindQuery(&queryParam); err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	user, err := u.service.List(c, &queryParam)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}
	var res []schema.User
	copier.Copy(&res, &user)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}
