package routes

import "github.com/gin-gonic/gin"

type GameType int

const (
	SauGras GameType = iota
	SauSchell
	SauEichel
	SoloHerz
	SoloGras
	SoloSchell
	SoloEichel
	Wenz
	Ramsch
)

type Game struct {
	Participants [4]int `json:"participants" binding:"required"`
	Player       int    `json:"player" binding:"required"`
	Playmate     int    `json:"playmate" binding:"required"`

	GameType GameType `json:"game_type" binding:"required"`
	Points   int      `json:"points" binding:"required"`
	Schwarz  bool     `json:"schwarz" binding:"required"`
}

// NewGame handles the /newGame route
func NewGame(c *gin.Context) {
	var g Game
	if c.ShouldBindJSON(&g) != nil {
		c.JSON(200, gin.H{
			"error": "invalid input format",
		})
	}

	// write game to db
}
