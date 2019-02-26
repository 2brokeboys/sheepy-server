package middleware

import "github.com/gin-gonic/gin"

// NoCache sets Cache-Control to "no-cache"
func NoCache(c *gin.Context) {
	c.Next()
	c.Header("Cache-Control", "no-cache")
}
