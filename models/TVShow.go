package models

type TVShow struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Desc        string  `json:"desc"`
	NumEpisodes int     `json:"numEpisodes"`
	Director    string  `json:"director"`
	Rating      float64 `json:"rating"`
	Genre       string  `json:"genre"`
}
