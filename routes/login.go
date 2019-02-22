package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Login serves the /login route
func Login(c *gin.Context) {
	session := sessions.Default(c)

	// check if already logged in
	v := session.Get("login")
	b, ok := v.(bool)
	if b && ok {
		c.JSON(200, gin.H{
			"error": "already logged in",
		})
		return
	}

	// get inputs
	var l struct {
		User     string `json:"user" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if c.ShouldBindJSON(&l) != nil {
		c.JSON(200, gin.H{
			"error": "inavlid data",
		})
		return
	}

	// validate credentials
	
	// update session
	session.Set("login", true)
	session.Set("uid", 23)
	session.Save()
	c.JSON(200, gin.H{"success": true})
}
