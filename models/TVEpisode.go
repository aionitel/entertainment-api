package models

type TVEpisode struct {
	ID      int64   `json:"id"`
	Title   string  `json:"title"`
	Desc    string  `json:"desc"`
	Season  int     `json:"season"`
	RunTime float32 `json:"runtime"`
	Rating  float64 `json:"rating"`
}
