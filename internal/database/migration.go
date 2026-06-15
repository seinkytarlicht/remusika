package database

import (
	"database/sql"
	"embed"
)

var migrationFiles embed.FS

func RunMigration(db *sql.DB) error {
	files := []string{}

	for _, file := range files {
		data, err := migrationFiles.ReadFile(file)

		if err != nil {
			return err
		}

		_, err = db.Exec(string(data))

		if err != nil {
			return err
		}
	}

	return nil
}
