package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Logout handles the /logout route
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user");
	session.Save();
	c.JSON(200, gin.H{
		"success": true,
	})
}
