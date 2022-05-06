package models

type Movie struct {
	ID       int64   `json:"id"`
	Title    string  `json:"title"`
	Desc     string  `json:"desc"`
	Director string  `json:"director"`
	Year     int     `json:"year"`
	Rating   float64 `json:"rating"`
}

func (movie *Movie) NewMovie(
	id int64,
	title string,
	desc string,
	director string,
	year int,
	rating float64,
) *Movie {
	return &Movie{
		ID:       id,
		Title:    title,
		Desc:     desc,
		Director: director,
		Year:     year,
		Rating:   rating,
	}
}
