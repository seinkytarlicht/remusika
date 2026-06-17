package service

import (
	"context"

	"github.com/seinkytarlicht/remusika/internal/model"
	"github.com/seinkytarlicht/remusika/internal/repository"
)

type PlaylistService interface {
	Create(ctx context.Context, playlist model.Playlist) (model.Playlist, error)
	Update(ctx context.Context, playlist model.Playlist) (model.Playlist, error)
	Delete(ctx context.Context, id uint64) bool
	FindAll(ctx context.Context) []model.Playlist
	AddMusic(ctx context.Context, playlistId uint64, Uuid string) error
	RemoveMusic(ctx context.Context, playlistItemId uint64) error
}

type PlaylistServiceImpl struct {
	PlaylistRepo repository.PlaylistRepository
	MusicMemo    repository.MusicMemo
}

func (p *PlaylistServiceImpl) Create(ctx context.Context, playlist model.Playlist) (model.Playlist, error) {
	result, err := p.PlaylistRepo.Create(ctx, playlist)

	return result, err
}

func (p *PlaylistServiceImpl) Update(ctx context.Context, playlist model.Playlist) (model.Playlist, error) {
	result, err := p.PlaylistRepo.Update(ctx, playlist)

	return result, err
}

func (p *PlaylistServiceImpl) Delete(ctx context.Context, id uint64) bool {
	result := p.PlaylistRepo.Delete(ctx, id)

	return result
}

func (p *PlaylistServiceImpl) FindAll(ctx context.Context) []model.Playlist {
	musics := p.MusicMemo.FindAll()
	playlists := p.PlaylistRepo.FindAll(ctx)

	musicMap := make(map[string]model.Music)
	for _, m := range musics {
		musicMap[m.Uuid] = m
	}

	for i := range playlists {
		if playlists[i].PlaylistItems == nil {
			playlists[i].PlaylistItems = []model.PlaylistItem{}
			continue
		}

		for j := range playlists[i].PlaylistItems {
			item := &playlists[i].PlaylistItems[j]

			if musicData, found := musicMap[item.MusicId]; found {
				item.Music = musicData
			}
		}
	}

	return playlists
}

func (p *PlaylistServiceImpl) Find(ctx context.Context, id uint64) model.Playlist {
	musics := p.MusicMemo.FindAll()
	playlist := p.PlaylistRepo.Find(ctx, id)

	musicMap := make(map[string]model.Music)
	for _, m := range musics {
		musicMap[m.Uuid] = m
	}

	if playlist.PlaylistItems == nil {
		playlist.PlaylistItems = []model.PlaylistItem{}
	}

	for j := range playlist.PlaylistItems {
		item := &playlist.PlaylistItems[j]

		if musicData, found := musicMap[item.MusicId]; found {
			item.Music = musicData
		}
	}

	return playlist
}

func (p *PlaylistServiceImpl) AddMusic(ctx context.Context, playlistId uint64, Uuid string) error {
	return p.PlaylistRepo.AddMusic(ctx, playlistId, Uuid)
}

func (p *PlaylistServiceImpl) RemoveMusic(ctx context.Context, playlistItemId uint64) error {
	return p.PlaylistRepo.RemoveMusic(ctx, playlistItemId)
}

func NewPlaylistService(playlistRepo repository.PlaylistRepository, musicMemo repository.MusicMemo) PlaylistService {
	return &PlaylistServiceImpl{
		PlaylistRepo: playlistRepo,
		MusicMemo:    musicMemo,
	}
}
