package model

type Movie struct {
	Title      string   `json:"title"`
	Year       string   `json:"year"`
	AgeRating  string   `json:"age_rating"`
	Runtime    string   `json:"runtime"`
	Genre      string   `json:"genre"`
	Director   string   `json:"director"`
	Writer     string   `json:"writer"`
	Cast       []string `json:"cast"`
	Plot       string   `json:"plot"`
	Awards     string   `json:"awards"`
	Poster     string   `json:"poster"`
	Ratings    []string `json:"ratings"`
	Metascore  string   `json:"metascore"`
	IMDBRating string   `json:"imdb_rating"`
}

func NewMovie( // to create new Movie object
	title string,
	year string,
	ageRating string,
	runtime string,
	genre string,
	director string,
	writer string,
	cast []string,
	plot string,
	awards string,
	poster string,
	ratings []string,
	metascore string,
	imdbRating string,
) *Movie {
	return &Movie{
		Title:      title,
		Year:       year,
		AgeRating:  ageRating,
		Runtime:    runtime,
		Genre:      genre,
		Director:   director,
		Writer:     writer,
		Cast:       cast,
		Plot:       plot,
		Awards:     awards,
		Poster:     poster,
		Ratings:    ratings,
		Metascore:  metascore,
		IMDBRating: imdbRating,
	}
}
