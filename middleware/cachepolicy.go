package middleware

import "github.com/gin-gonic/gin"

// CacheControl sets Cache-Control header
func CacheControl(c *gin.Context) {
	c.Header("Cache-Control", "public, max-age=31536000")
}
