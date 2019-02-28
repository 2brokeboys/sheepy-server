package routes

import (
	"github.com/2brokeboys/sheepy-server/common"
	"github.com/2brokeboys/sheepy-server/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Login serves the /login route
func Login(c *gin.Context) {
	session := sessions.Default(c)

	// Check if already logged in
	v := session.Get("user")
	_, ok := v.(*common.User)
	if ok {
		c.JSON(200, gin.H{
			"error": "already logged in",
		})
		return
	}

	// Get inputs
	var l struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := c.ShouldBindJSON(&l)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "invalid data",
		})
		return
	}

	// Validate credentials
	user, ok := db.AuthentificateUser(l.Username, l.Password)
	if !ok {
		c.JSON(200, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	// Update session
	session.Set("user", user)
	session.Save()
	c.JSON(200, gin.H{
		"success": true,
		"user":    user,
	})
}
