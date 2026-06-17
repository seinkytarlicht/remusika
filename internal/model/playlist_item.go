package model

type PlaylistItem struct {
	Id         uint64 `json:"id" db:"Id"`
	PlaylistId uint64 `json:"playlist_id" db:"PlaylistId"`
	MusicId    string `json:"music_id" db:"MusicId"`
	OrderPos   string `json:"order_pos" db:"OrderPos"`
	Music      Music  `json:"music"`
}
