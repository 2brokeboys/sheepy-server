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
	Participants [4]int `json:"participants"`
	Player       int    `json:"player"`
	Playmate     int    `json:"playmate"`

	// Game information
	GameType GameType `json:"gameType"`
	Points   int      `json:"points"`
	Schwarz  bool     `json:"schwarz"`
	Runners  int      `json:"runners"`
	Virgins  int      `json:"virgins"`

	// Meta information
	Time     time.Time `json:"-"`
	Reporter int       `json:"-"`
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
