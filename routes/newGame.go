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
	if c.ShouldBindJSON(g) != nil {
		c.JSON(400, gin.H{
			"error": "invalid data",
		})
	}

	g.Reporter = c.MustGet("user").(*common.User).ID
	g.Time = time.Now()

	// Do input validation
	//FIXME

	// Write game to database
	db.InsertGame(g)
}
