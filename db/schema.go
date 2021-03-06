package db

import (
	"time"

	"github.com/2brokeboys/sheepy-server/common"
	"github.com/pkg/errors"
)

const schema = `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username varchar[16] UNIQUE,
	name varchar[100],
	pw varchar[60]
);

CREATE TABLE IF NOT EXISTS games (
	part0 int,
	part1 int,
	part2 int,
	part3 int,

	player int,
	playmate int,

	gametype tiny,
	points tiny,
	schwarz bit,
	runners tiny,
	virgins tiny,

	time datetime,
	reporter int
);`

type dbUser struct {
	ID       int
	Username string
	Name     string
	Pw       string
}

type dbGame struct {
	Part0 int
	Part1 int
	Part2 int
	Part3 int

	Player   int
	Playmate int

	Gametype int8
	Points   int8
	Schwarz  bool
	Runners  int8
	Virgins  int8

	Time     time.Time
	Reporter int
}

func migrate() error {
	_, err := db.Exec(schema)
	return errors.Wrap(err, "error migrating")
}

func (g *dbGame) indexInParticipants(uid int) int {
	switch uid {
	case g.Part0:
		return 0
	case g.Part1:
		return 1
	case g.Part2:
		return 2
	case g.Part3:
		return 3
	}
	// this shouldn't happen
	return -1
}

func (user *dbUser) ToCommon() *common.User {
	return &common.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
}

func (g *dbGame) ToCommon() *common.Game {
	return &common.Game{
		Participants: [4]int{g.Part0, g.Part1, g.Part2, g.Part3},
		Player:       g.indexInParticipants(g.Player),
		Playmate:     g.indexInParticipants(g.Playmate),

		GameType: common.GameType(g.Gametype),
		Points:   int(g.Points),
		Schwarz:  g.Schwarz,
		Runners:  int(g.Runners),
		Virgins:  int(g.Virgins),

		Time:     g.Time,
		Reporter: g.Reporter,
	}
}
