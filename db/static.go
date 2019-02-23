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
	db, err = sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		return errors.Wrap(err, "Error connecting to database")
	}
	err = migrate()
	if err != nil {
		return errors.Wrap(err, "Error migrating schema")
	}
	err = initStatements()
	if err != nil {
		return errors.Wrap(err, "Error initing Statements")
	}
	return nil
}
