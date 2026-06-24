package model

type PlaylistItem struct {
	Id         uint64 `json:"id" db:"Id"`
	PlaylistId uint64 `json:"playlist_id" db:"PlaylistId"`
	MusicId    string `json:"music_id" db:"MusicId"`
	Pos        int64  `json:"order" db:"Pos"`
	CreatedAt  string `json:"created_at" db:"CreatedAt"`
	Music      Music  `json:"music"`
}
