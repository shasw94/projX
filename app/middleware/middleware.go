package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// SkipperFunc skipper function
type SkipperFunc func(*gin.Context) bool

// AllowPathPrefixSkipper allow paths have prefix skip middleware
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)
		for _, p := range prefixes {
			if plen := len(p); pathLen >= plen && path[:plen] == p {
				return true
			}
		}
		return false
	}
}

// AllowPathPrefixNoSkipper allow paths have prefixes no skip middleware
func AllowPathPrefixNoSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}

func JoinRouter(method, path string) string {
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(method), path)
}

// SkipHandler skip handler
func SkipHandler(c *gin.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(c) {
			return true
		}
	}
	return false
}

// EmptyMiddleware return empty middleware
func EmptyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}