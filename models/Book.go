package models

type Book struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Desc   string  `json:"desc"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Rating float64 `json:"rating"`
}
