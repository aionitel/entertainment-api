package models

type Movie struct {
	ID       int64   `json:"id"`
	Title    string  `json:"title"`
	Desc     string  `json:"desc"`
	Director string  `json:"director"`
	Year     int     `json:"year"`
	Rating   float64 `json:"rating"`
}
