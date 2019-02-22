package common

import (
	"encoding/gob"
	"time"
)

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

// Game represents one game played
type Game struct {
	// Player information
	Participants [4]int `json:"participants" binding:"required"`
	Player       int    `json:"player" binding:"required"`
	Playmate     int    `json:"playmate" binding:"required"`

	// Game information
	GameType GameType `json:"game_type" binding:"required"`
	Points   int      `json:"points" binding:"required"`
	Schwarz  bool     `json:"schwarz" binding:"required"`

	// Meta information
	Time     time.Time `json:"time"`
	Reporter int
}

// User represents a user of the website
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

func init() {
	gob.Register(&Game{})
	gob.Register(&User{})
}
