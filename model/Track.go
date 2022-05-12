package model

type Album struct {
	Title       string `json:"title"`
	CoverSmall  string `json:"cover_small"`
	CoverMedium string `json:"cover_medium"`
	CoverBig    string `json:"cover_big"`
	CoverXl     string `json:"cover_xl"`
	Tracklist   string `json:"tracklist"`
	Type        string `json:"type"`
}

type Artist struct {
	Name          string `json:"name"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	Type          string `json:"type"`
}

type Track struct {
	Title          string `json:"title"`
	Duration       int    `json:"duration"`
	Rank           int    `json:"rank"`
	ExplicitLyrics bool   `json:"explicit_lyrics"`
	Preview        string `json:"preview"`
	Artist         Artist `json:"artist"`
	Album          Album  `json:"album"`
}

type Music struct {
	Tracks []Track `json:"data"`
}
