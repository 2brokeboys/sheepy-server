package routes

import "github.com/gin-gonic/gin"

// Root handles the / route by serving index.html
func Root(c *gin.Context) {
	// FIXME: serve index.html here and server-push required assets
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
