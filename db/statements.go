package db

import (
	"log"
	"strings"
	"sync"

	"github.com/2brokeboys/sheepy-server/common"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var (
	getExactUserStatement     *sqlx.Stmt
	insertGameStatement       *sqlx.NamedStmt
	queryUserStatement        *sqlx.Stmt
	getUserStatement          *sqlx.Stmt
	insertUserStatement       *sqlx.NamedStmt
	queryRecentGamesStatement *sqlx.Stmt
)

func initStatements() error {
	var err error

	getExactUserStatement, err = db.Preparex("SELECT * FROM users WHERE LOWER(username)=LOWER(?)")
	if err != nil {
		return errors.Wrap(err, "Error preparing getExactUserStatement")
	}

	insertGameStatement, err = db.PrepareNamed(`INSERT INTO games (part0, part1, part2, part3, player, playmate, gametype, points, schwarz, runners, virgins, time, reporter)
	VALUES (:part0, :part1, :part2, :part3, :player, :playmate, :gametype, :points, :schwarz, :runners, :virgins, :time, :reporter)`)
	if err != nil {
		return errors.Wrap(err, "Error preparing insertGameStatement")
	}

	queryUserStatement, err = db.Preparex("SELECT * FROM users WHERE LOWER(username) LIKE ? OR LOWER(name) LIKE ? LIMIT 20")
	if err != nil {
		return err
	}

	insertUserStatement, err = db.PrepareNamed("INSERT INTO users (username, name, pw) VALUES (:username, :name, :pw)")
	if err != nil {
		return err
	}

	queryRecentGamesStatement, err = db.Preparex("SELECT * FROM games LIMIT ?, ?")
	if err != nil {
		return err
	}

	return nil
}

// InsertGame inserts the given game into the db
func InsertGame(game *common.Game) error {
	log.Println("player", game.Player)
	playmate := -1
	if game.Playmate != -1 {
		playmate = game.Participants[game.Playmate]
	}
	log.Println("playmate", game.Playmate)
	dbgame := &dbGame{
		Part0: game.Participants[0],
		Part1: game.Participants[1],
		Part2: game.Participants[2],
		Part3: game.Participants[3],

		Player:   game.Participants[game.Player],
		Playmate: playmate,

		Gametype: int8(game.GameType),
		Points:   int8(game.Points),
		Schwarz:  game.Schwarz,
		Virgins:  int8(game.Virgins),
		Runners:  int8(game.Runners),

		Time:     game.Time,
		Reporter: game.Reporter,
	}

	_, err := insertGameStatement.Exec(dbgame)
	if err != nil {
		return errors.Wrap(err, "Error inserting game into db")
	}

	return nil
}

// Apparently there is some bug which hinders concurrency
var mu sync.Mutex

// QueryUser returns all users matching the given search string
func QueryUser(search string) ([]*common.User, error) {
	mu.Lock()
	defer mu.Unlock()
	dbUsers := make([]dbUser, 0)

	// simulate half-fuzzy search
	search = strings.ToLower(search)
	search = strings.Replace(search, " ", "%", -1)
	search = "%" + search + "%"

	err := queryUserStatement.Select(&dbUsers, search, search)
	if err != nil {
		return nil, errors.Wrap(err, "Error querying users")
	}
	// FIXME: sort by quality of match?
	ret := make([]*common.User, len(dbUsers))
	for i := range ret {
		ret[i] = dbUsers[i].ToCommon()
	}
	return ret, nil
}

// GetUser returns all users matching the given search string
// Returns nil when the user was not found (no error)
func GetUser(username string) (*common.User, error) {
	dbUser := dbUser{}
	err := getExactUserStatement.Get(&dbUser, username)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			// If no user was found, this is not an error
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error querying users")
	}

	return dbUser.ToCommon(), nil
}

// QueryRecentGames return the numer recent games
func QueryRecentGames(from, number int) ([]*common.Game, error) {
	dbGames := make([]dbGame, 0)
	err := queryRecentGamesStatement.Select(&dbGames, from, number)
	if err != nil {
		return nil, errors.Wrap(err, "Error querying games")
	}
	ret := make([]*common.Game, len(dbGames))
	for i := range ret {
		ret[i] = dbGames[i].ToCommon()
	}
	return ret, nil
}
