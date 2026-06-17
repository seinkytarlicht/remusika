package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/seinkytarlicht/remusika/internal/model"
)

type PlaylistRepository interface {
	Create(ctx context.Context, playlist model.Playlist) (model.Playlist, error)
	Update(ctx context.Context, playlist model.Playlist) (model.Playlist, error)
	Delete(ctx context.Context, id uint64) bool
	Find(ctx context.Context, id uint64) model.Playlist
	FindAll(ctx context.Context) []model.Playlist
	AddMusic(ctx context.Context, playlistId uint64, Uuid string) error
	RemoveMusic(ctx context.Context, playlistItemId uint64) error
}

type PlaylistRepositoryImpl struct {
	DB *sqlx.DB
}

func (p *PlaylistRepositoryImpl) Create(ctx context.Context, playlist model.Playlist) (model.Playlist, error) {
	isExist := model.Playlist{}

	err := p.DB.GetContext(ctx, &isExist, "SELECT * FROM Playlist WHERE Name=$1", playlist.Name)
	if err == nil {
		return isExist, errors.New("Playlist name already exist")
	}

	res, err := p.DB.NamedExecContext(ctx, "INSERT INTO Playlist (Name) VALUES (:Name)", playlist)

	if err != nil {
		panic(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return p.Find(ctx, uint64(lastId)), nil
}

func (p *PlaylistRepositoryImpl) Update(ctx context.Context, playlist model.Playlist) (model.Playlist, error) {
	isExist := model.Playlist{}

	err := p.DB.GetContext(ctx, &isExist, "SELECT * FROM Playlist WHERE Name=$1 AND NOT Id=$2", playlist.Name, playlist.Id)
	if err == nil {
		return isExist, errors.New("Playlist name already exist")
	}

	_, err = p.DB.NamedExecContext(ctx, "UPDATE Playlist SET Name=:Name WHERE Id=:Id", playlist)

	if err != nil {
		panic(err)
	}

	return p.Find(ctx, playlist.Id), nil
}

func (p *PlaylistRepositoryImpl) Delete(ctx context.Context, id uint64) bool {
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

	playlistMap := make(map[uint64]*model.Playlist)

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

func (p *PlaylistRepositoryImpl) Find(ctx context.Context, id uint64) model.Playlist {
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

func (p *PlaylistRepositoryImpl) AddMusic(ctx context.Context, playlistId uint64, Uuid string) error {
	isExist := model.PlaylistItem{}

	err := p.DB.GetContext(ctx, &isExist, "SELECT * FROM PlaylistItem WHERE PlaylistId=$1 AND MusicId=$2", playlistId, Uuid)
	if err == nil {
		return errors.New("Music already in playlist")
	}

	res, err := p.DB.NamedExecContext(ctx, "INSERT INTO PlaylistItem (PlaylistId, MusicId) VALUES (:playlist_id, :music_id)", map[string]any{
		"playlist_id": playlistId,
		"music_id":    Uuid,
	})

	if err != nil {
		panic(err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return nil
}

func (p *PlaylistRepositoryImpl) RemoveMusic(ctx context.Context, playlistItemId uint64) error {
	_, err := p.DB.NamedExecContext(ctx, "DELETE FROM PlaylistItem WHERE Id=:id", map[string]any{
		"id": playlistItemId,
	})

	if err != nil {
		return errors.New("Failed to remove music from playlist")
	}

	return nil
}

func NewPlaylistRepository(db *sqlx.DB) PlaylistRepository {
	return &PlaylistRepositoryImpl{
		DB: db,
	}
}
