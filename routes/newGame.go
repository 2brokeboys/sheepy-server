package routes

import (
	"time"

	"github.com/2brokeboys/sheepy-server/common"
	"github.com/2brokeboys/sheepy-server/db"
	"github.com/gin-gonic/gin"
)

// NewGame handles the /newGame route
func NewGame(c *gin.Context) {
	g := &common.Game{}
	err := c.ShouldBindJSON(g)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid data",
		})
		return
	}

	g.Reporter = c.MustGet("user").(*common.User).ID
	g.Time = time.Now()

	// Do input validation
	msgs := g.Sanitize()
	if len(msgs) > 0 {
		c.JSON(400, gin.H{
			"error":    "invalid game",
			"messages": msgs,
		})
	}

	// Write game to database
	db.InsertGame(g)

	c.JSON(200, gin.H{
		"success": true,
	})
}
