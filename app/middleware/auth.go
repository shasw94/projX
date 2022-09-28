package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shasw94/projX/pkg/app"
	"github.com/shasw94/projX/pkg/http/wrapper"
	"github.com/shasw94/projX/pkg/jwt"
)

func wrapUserAuthContext(c *gin.Context, userId string) {
	app.SetUserID(c, userId)
	c.Request = c.Request.WithContext(c)
}

// UserAuthMiddleware User Auth Middleware
func UserAuthMiddleware(a jwt.IJWTAuth, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userID, err := a.ParseUserID(app.GetToken(c), false)
		if err != nil {
			wrapper.Translate(c, wrapper.Response{Error: err})
			c.Abort()
			return
		}
		wrapUserAuthContext(c, userID)
		c.Next()
	}
}
