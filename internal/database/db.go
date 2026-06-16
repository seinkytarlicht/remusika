package database

import (
	_ "modernc.org/sqlite"

	"github.com/jmoiron/sqlx"
	"github.com/seinkytarlicht/remusika/config"
)

func New() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite", config.GetAppDBFile())

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	// if err := RunMigration(db); err != nil {
	// 	return nil, err
	// }

	return db, nil
}
