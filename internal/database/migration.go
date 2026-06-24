package database

import (
	"embed"
	"path/filepath"

	"github.com/jmoiron/sqlx"
)

//go:embed all:migrations
var migrationFiles embed.FS

func RunMigration(db *sqlx.DB, skip int64) error {
	files := []string{
		"001_remusika.sql",
	}
	sql := "INSERT INTO Migration (LastUpdate) VALUES (:li)"
	var lastUpdate int64
	var lastIndex int64

	if skip != -1 {
		lastUpdate = skip
		files = files[skip:]
		sql = "UPDATE Migration SET LastUpdate=:li WHERE Id=1"
	}

	for i, file := range files {
		data, err := migrationFiles.ReadFile(filepath.Join("migrations", file))

		if err != nil {
			return err
		}

		_, err = db.Exec(string(data))

		if err != nil {
			return err
		}

		lastIndex = int64(i + 1)
	}

	lastUpdate += lastIndex

	db.NamedExec(sql, map[string]any{
		"li": lastUpdate,
	})

	return nil
}
