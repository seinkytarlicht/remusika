package model

import "database/sql"

type Playlist struct {
	Id            uint64         `json:"id" db:"Id"`
	Name          string         `json:"name" db:"Name"`
	Description   sql.NullString `json:"description" db:"Description"`
	Image         sql.NullString `json:"image" db:"Image"`
	Pos           int64          `json:"order" db:"Pos"`
	CreatedAt     string         `json:"created_at" db:"CreatedAt"`
	PlaylistItems []PlaylistItem `json:"playlist_items"`
}
