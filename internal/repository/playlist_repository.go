package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/seinkytarlicht/remusika/internal/model"
)

type PlaylistRepository interface {
	Create(ctx context.Context, playlist model.Playlist) model.Playlist
	Update(ctx context.Context, playlist model.Playlist) model.Playlist
	Delete(ctx context.Context, id int64) bool
	Find(ctx context.Context, id int64) model.Playlist
	FindAll(ctx context.Context) []model.Playlist
}

type PlaylistRepositoryImpl struct {
	DB *sqlx.DB
}

func (p *PlaylistRepositoryImpl) Create(ctx context.Context, playlist model.Playlist) model.Playlist {
	res, err := p.DB.NamedExecContext(ctx, "INSERT INTO Playlist (name) VALUES (:Name)", playlist)

	if err != nil {
		panic(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return p.Find(ctx, lastId)
}

func (p *PlaylistRepositoryImpl) Update(ctx context.Context, playlist model.Playlist) model.Playlist {
	_, err := p.DB.NamedExecContext(ctx, "UPDATE Playlist SET Name=:Name WHERE Id=:Id", playlist)

	if err != nil {
		panic(err)
	}

	return p.Find(ctx, int64(playlist.Id))
}

func (p *PlaylistRepositoryImpl) Delete(ctx context.Context, id int64) bool {
	_, err := p.DB.NamedExecContext(ctx, "DELETE FROM Playlist WHERE Id=:id", map[string]any{
		"id": id,
	})

	if err != nil {
		return false
	}

	return true
}

func (p *PlaylistRepositoryImpl) FindAll(ctx context.Context) []model.Playlist {
	playlists := []model.Playlist{}
	playlistItems := []model.PlaylistItem{}

	err := p.DB.SelectContext(ctx, &playlists, "SELECT * FROM Playlist")
	if err != nil {
		panic(err)
	}

	err = p.DB.SelectContext(ctx, &playlistItems, "SELECT * FROM PlaylistItem")
	if err != nil {
		panic(err)
	}

	playlistMap := make(map[int64]*model.Playlist)

	for i := range playlists {
		playlistMap[playlists[i].Id] = &playlists[i]
	}

	for _, pItem := range playlistItems {
		playlist := playlistMap[pItem.PlaylistId]

		if playlist != nil {
			playlist.PlaylistItems = append(playlist.PlaylistItems, pItem)
		}
	}

	return playlists
}

func (p *PlaylistRepositoryImpl) Find(ctx context.Context, id int64) model.Playlist {
	playlist := model.Playlist{}
	playlistItems := []model.PlaylistItem{}

	err := p.DB.GetContext(ctx, &playlist, "SELECT * FROM Playlist WHERE Id=$1", id)
	if err != nil {
		panic(err)
	}

	err = p.DB.SelectContext(ctx, &playlistItems, "SELECT * FROM PlaylistItem WHERE PlaylistId=$1", id)
	if err != nil {
		panic(err)
	}

	playlist.PlaylistItems = playlistItems

	return playlist
}

func NewPlaylistRepository(db *sqlx.DB) PlaylistRepository {
	return &PlaylistRepositoryImpl{
		DB: db,
	}
}
