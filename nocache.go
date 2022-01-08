package wnocache

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var expires = time.Unix(0, 0).UTC().Format(http.TimeFormat)

// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Cache-Control
var responseHeaders = map[string]string{
	"Cache-Control":   "no-cache, no-store, max-age=0, must-revalidate, no-transform",
	"Pragma":          "no-cache",
	"Expires":         expires,
	"X-Accel-Expires": "0", // nginx
}

var etagHeaders = []string{
	"ETag",
	"If-Modified-Since",
	"If-Match",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
}

// New removes ETag related headers and injects regular "no cache" headers
func New() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, value := range etagHeaders {
			if c.Request.Header.Get(value) != "" {
				c.Request.Header.Del(value)
			}
		}

		for key, value := range responseHeaders {
			c.Writer.Header().Set(key, value)
		}
		c.Next()
	}
}
