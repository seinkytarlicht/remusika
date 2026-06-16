package service

import (
	"context"

	"github.com/seinkytarlicht/remusika/internal/model"
	"github.com/seinkytarlicht/remusika/internal/repository"
)

type PlaylistService interface {
	Create(ctx context.Context, playlist model.Playlist) model.Playlist
	Update(ctx context.Context, playlist model.Playlist) model.Playlist
	Delete(ctx context.Context, id int64) bool
	FindAll(ctx context.Context) []model.Playlist
}

type PlaylistServiceImpl struct {
	Repo repository.PlaylistRepository
}

func (p *PlaylistServiceImpl) Create(ctx context.Context, playlist model.Playlist) model.Playlist {
	result := p.Repo.Create(ctx, playlist)

	return result
}

func (p *PlaylistServiceImpl) Update(ctx context.Context, playlist model.Playlist) model.Playlist {
	result := p.Repo.Update(ctx, playlist)

	return result
}

func (p *PlaylistServiceImpl) Delete(ctx context.Context, id int64) bool {
	result := p.Repo.Delete(ctx, id)

	return result
}

func (p *PlaylistServiceImpl) FindAll(ctx context.Context) []model.Playlist {
	return p.Repo.FindAll(ctx)
}

func NewPlaylistService(repo repository.PlaylistRepository) PlaylistService {
	return &PlaylistServiceImpl{
		Repo: repo,
	}
}
