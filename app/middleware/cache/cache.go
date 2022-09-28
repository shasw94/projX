package cache

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/shasw94/projX/logger"
	"net/http"
	"strings"
)

type Cache interface {
	IsConnected() bool
	Get(key string, data interface{}) error
	Set(key string, val []byte) error
	Remove(keys ...string) error
	Keys(pattern string) ([]string, error)
}

// New Setup Initialize the Cache instance
func New() Cache {
	return NewRedis()
}

var cache = New()

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Cached middleware
func Cached() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cache == nil || !cache.IsConnected() {
			logger.Warn("Cache is not available")
			c.Next()
			return
		}

		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		key := c.Request.URL.RequestURI()
		if c.Request.Method != "GET" {
			c.Next()

			statusCode := w.Status()
			if statusCode != http.StatusOK {
				return
			}

			if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
				temp := strings.Split(key, "/")
				objName := temp[len(temp)-1]
				keys, _ := cache.Keys("*" + objName + "*")
				if keys != nil {
					cache.Remove(keys...)
				}
			}
			return
		}

		var data map[string]interface{}
		cache.Get(key, &data)

		if data != nil {
			c.JSON(http.StatusOK, data)
			c.Abort()
			return
		}
		c.Next()
		statusCode := w.Status()
		if statusCode == http.StatusOK {
			cache.Set(key, w.body.Bytes())
		}
	}
}
