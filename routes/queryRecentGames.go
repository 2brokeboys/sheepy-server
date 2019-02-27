package routes

import (
	"log"

	"github.com/2brokeboys/sheepy-server/db"
	"github.com/gin-gonic/gin"
)

// QueryRecentGames handles the /queryRecentGames route
func QueryRecentGames(c *gin.Context) {
	var p struct {
		From   int `json:"from"`
		Number int `json:"number" binding:"required"`
	}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "invalid data",
		})
		return
	}

	if p.From < 0 {
		c.JSON(200, gin.H{
			"error": "index out of range",
		})
		return
	}
	if p.Number < 0 || p.Number > 50 {
		c.JSON(200, gin.H{
			"error": "number has to be within 0 to 50",
		})
		return
	}

	games, err := db.QueryRecentGames(p.From, p.Number)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"error": "database error",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"games":   games,
	})
}
