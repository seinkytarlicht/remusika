package database

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"github.com/seinkytarlicht/remusika/config"
)

func New() (*sql.DB, error) {
	db, err := sql.Open("sqlite", config.GetAppDBFile())

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	// if err := RunMigration(db); err != nil {
	// 	return nil, err
	// }

	return db, nil
}
