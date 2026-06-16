package model

type PlaylistItem struct {
	Id         int64  `json:"id" db:"Id"`
	PlaylistId int64  `json:"playlist_id" db:"PlaylistId"`
	MusicId    string `json:"music_id" db:"MusicId"`
	OrderPos   string `json:"order" db:"OrderPos"`
}
