package repository

import (
	"context"
	"database/sql"
)

type PlaylistRepository interface {
	Exec(ctx context.Context, query string, args ...any) (sql.Result, error)
	Query(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type PlaylistRepositoryImpl struct {
	DB *sql.DB
}

func (p *PlaylistRepositoryImpl) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		panic(err)
	}

	result, err := tx.ExecContext(ctx, query, args...)

	return result, err
}

func (p *PlaylistRepositoryImpl) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		panic(err)
	}

	rows, err := tx.QueryContext(ctx, query, args...)

	return rows, err
}

func NewPlaylistRepository(db *sql.DB) PlaylistRepository {
	return &PlaylistRepositoryImpl{
		DB: db,
	}
}
