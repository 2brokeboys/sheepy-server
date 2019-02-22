package db

import "github.com/pkg/errors"

const schema = `CREATE TABLE IF NOT EXISTS users (
	username varchar[60],
	pw varchar[100]
);`

func migrate() error {
	_, err := db.Exec(schema)
	return errors.Wrap(err, "error migrating")
}
