package routes

import (
	"github.com/2brokeboys/sheepy-server/db"
	"github.com/gin-gonic/gin"
)

// GetUser handles the /queryUser route
func GetUser(c *gin.Context) {
	var p struct {
		Username string `json:"username" binding:"required"`
	}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "invalid data",
		})
		return
	}

	user, err := db.GetUser(p.Username)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "error querying database",
		})
		return
	}
	if user == nil {
		c.JSON(200, gin.H{
			"success": true,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"user":    user,
	})
}
