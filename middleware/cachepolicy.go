package middleware

import "github.com/gin-gonic/gin"

// CacheControl sets Cache-Control header
func CacheControl(c *gin.Context) {
	if c.Request.URL.RequestURI() == "/" {
		return
	}
	c.Header("Cache-Control", "public, max-age=31536000")
}
