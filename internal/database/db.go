package database

import (
	_ "modernc.org/sqlite"

	"github.com/jmoiron/sqlx"
	"github.com/seinkytarlicht/remusika/config"
	"github.com/seinkytarlicht/remusika/internal/model"
)

func New() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite", config.GetAppDBFile())

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	migration := model.Migration{}

	if err := db.Get(&migration, "SELECT * FROM migration WHERE Id=1"); err != nil {
		if err := RunMigration(db, -1); err != nil {
			return nil, err
		}
	}

	if err := RunMigration(db, migration.LastUpdate); err != nil {
		return nil, err
	}

	return db, nil
}
