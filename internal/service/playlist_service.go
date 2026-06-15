package service

import (
	"context"

	"github.com/seinkytarlicht/remusika/internal/model"
	"github.com/seinkytarlicht/remusika/internal/repository"
)

type PlaylistService interface {
	Create(ctx context.Context, playlist model.Playlist)
	FindAll(ctx context.Context) []model.Playlist
}

type PlaylistServiceImpl struct {
	Repo repository.PlaylistRepository
}

func (p *PlaylistServiceImpl) Create(ctx context.Context, playlist model.Playlist) {
	SQL := "INSERT INTO Playlists (name) VALUES (?);"

	p.Repo.Query(ctx, SQL, playlist.Name)
}

func (p *PlaylistServiceImpl) FindAll(ctx context.Context) []model.Playlist {
	SQL := "SELECT * FROM Playlist;"

	rows, err := p.Repo.Query(ctx, SQL)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var playlists []model.Playlist

	if rows.Next() {
		playlist := model.Playlist{}
		err := rows.Scan(&playlist.Id, &playlist.Name, &playlist.CreatedAt)
		if err == nil {
			playlists = append(playlists, playlist)
		}
	}

	return playlists
}

func NewPlaylistService(repo repository.PlaylistRepository) PlaylistService {
	return &PlaylistServiceImpl{
		Repo: repo,
	}
}
