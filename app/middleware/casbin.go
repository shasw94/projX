package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/shasw94/projX/app/contextx"
	"github.com/shasw94/projX/config"
	"github.com/shasw94/projX/pkg/utils"
	"net/http"
)

// CasbinMiddleware Valid user interface permission
func CasbinMiddleware(enforcer *casbin.SyncedEnforcer, skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := config.Config.Casbin
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		p := c.Request.URL.Path
		m := c.Request.Method
		userID := contextx.FromUserID(c.Request.Context())
		if b, err := enforcer.Enforce(userID, p, m); err != nil {
			c.JSON(http.StatusUnauthorized, utils.PrepareResponse(nil, "Unauthorized", utils.ErrorAuth))

			c.Abort()
			return
		} else if !b {
			c.JSON(http.StatusUnauthorized, utils.PrepareResponse(nil, "Unauthorized", utils.ErrorAuth))

			c.Abort()
			return
		}
		c.Next()
	}
}
