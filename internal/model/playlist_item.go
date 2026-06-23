package model

type PlaylistItem struct {
	Id         uint64 `json:"id" db:"Id"`
	PlaylistId uint64 `json:"playlist_id" db:"PlaylistId"`
	MusicId    string `json:"music_id" db:"MusicId"`
	Pos        string `json:"order" db:"Pos"`
	Music      Music  `json:"music"`
}
