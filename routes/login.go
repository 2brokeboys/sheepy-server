package routes

import (
	"log"

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
		c.JSON(409, gin.H{
			"error": "already logged in",
		})
		return
	}

	// Get inputs
	var l struct {
		User     string `json:"user" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := c.ShouldBindJSON(&l)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "invalid data",
		})
		return
	}

	// Validate credentials
	user, ok := db.AuthentificateUser(l.User, l.Password)
	if !ok {
		c.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	// Update session
	session.Set("user", user)
	session.Save()
	c.JSON(200, gin.H{
		"success": true,
	})
}
