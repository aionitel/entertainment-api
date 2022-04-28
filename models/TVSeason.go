package models

type TVSeason struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Desc        string      `json:"desc"`
	Seasons     int         `json:"seasons"`
	Rating      float64     `json:"rating"`
	NumEpisodes int         `json:"numEpisodes"`
	Episodes    []TVEpisode `json:"episodes"`
}
