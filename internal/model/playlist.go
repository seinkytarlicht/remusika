package model

type Playlist struct {
	Id            uint64         `json:"id" db:"Id"`
	Name          string         `json:"name" db:"Name"`
	CreatedAt     string         `json:"created_at" db:"CreatedAt"`
	PlaylistItems []PlaylistItem `json:"playlist_items"`
}
