package routes

import (
	"github.com/2brokeboys/sheepy-server/db"
	"github.com/gin-gonic/gin"
)

// QueryUser handles the /queryUser route
func QueryUser(c *gin.Context) {
	var p struct {
		Search string `json:"search" binding:"required"`
	}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "invalid data",
		})
		return
	}

	users, err := db.QueryUser(p.Search)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "error querying database",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"users":   users,
	})
}
