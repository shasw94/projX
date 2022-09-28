package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shasw94/projX/app/api"
	"github.com/shasw94/projX/app/middleware"
	"github.com/shasw94/projX/logger"
	"github.com/shasw94/projX/pkg/http/wrapper"
	"github.com/shasw94/projX/pkg/jwt"
	"go.uber.org/dig"
)

// RegisterAPI register api routes
func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		jwt jwt.IJWTAuth,
		authAPI *api.AuthAPI,
		userAPI *api.UserAPI,
		roleAPI *api.RoleAPI,
	) error {
		jwtMiddle := middleware.UserAuthMiddleware(jwt)
		//corsMiddle := middleware.CORSMiddleware()
		//casbinMiddle := middleware.CasbinMiddleware(casbinEnforcer)

		{
			r.POST("/register", wrapper.Wrap(authAPI.Register))
			r.POST("/login", wrapper.Wrap(authAPI.Login))
			r.POST("/refresh", wrapper.Wrap(authAPI.Refresh))
			r.POST("/logout", jwtMiddle, wrapper.Wrap(authAPI.Logout))
		}

		adminPath := r.Group("/admin", jwtMiddle)
		{
			adminPath.POST("/roles", wrapper.Wrap(roleAPI.CreateRole))
		}

		//-------------------------API---------------------------
		apiPath := r.Group("/api/v1", jwtMiddle)
		{
			apiPath.GET("/users/:id", userAPI.GetByID)
			apiPath.GET("/users", wrapper.Wrap(userAPI.List))
		}
		return nil
	})
	if err != nil {
		logger.Error(err)
	}

	return err
}
