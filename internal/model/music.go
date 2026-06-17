package model

type Music struct {
	Uuid        string `json:"uuid"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	ImageUrl    string `json:"image_url"`
	AudioUrl    string `json:"audio_url"`
	TempImgPath string `json:"-"`
	Path        string `json:"-"`
}
