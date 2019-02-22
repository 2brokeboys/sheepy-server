package db

import (
	// sqlite driver
	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var db *sqlx.DB

// InitDB initializes the database
func InitDB() error {
	var err error
	//db = sqlx.Connect("mysql", "root:root@localhost/db")
	db, err = sqlx.Connect("sqlite3", "/tmp/sheepy.sqlite3")
	if err != nil {
		return errors.Wrap(err, "Error connecting to database")
	}
	err = initStatements()
	if err != nil {
		return err
	}
	return migrate()
}
