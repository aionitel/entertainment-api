package model

type TV struct {
	Title        string   `json:"Title"`
	Year         string   `json:"Year"`
	AgeRating    string   `json:"Rated"`
	Runtime      string   `json:"Runtime"`
	Genre        string   `json:"Genre"`
	Director     string   `json:"Director"`
	Writer       string   `json:"Writer"`
	Cast         string   `json:"Actors"`
	Plot         string   `json:"Plot"`
	Awards       string   `json:"Awards"`
	Poster       string   `json:"Poster"`
	Ratings      []Rating `json:"Ratings"`
	Metascore    string   `json:"Metascore"`
	IMDBRating   string   `json:"imdbRating"`
	DVDRelease   string   `json:"DVD"`
	BoxOffice    string   `json:"BoxOffice"`
	TotalSeasons string   `json:"totalSeasons"`
}

func NewTV( // to create new Movie object
	title string,
	year string,
	ageRating string,
	runtime string,
	genre string,
	director string,
	writer string,
	cast string,
	plot string,
	awards string,
	poster string,
	ratings []Rating,
	metascore string,
	imdbRating string,
	dvdRelease string,
	boxOffice string,
	totalSeasons string,
) *TV {
	return &TV{
		Title:        title,
		Year:         year,
		AgeRating:    ageRating,
		Runtime:      runtime,
		Genre:        genre,
		Director:     director,
		Writer:       writer,
		Cast:         cast,
		Plot:         plot,
		Awards:       awards,
		Poster:       poster,
		Ratings:      ratings,
		Metascore:    metascore,
		IMDBRating:   imdbRating,
		DVDRelease:   dvdRelease,
		BoxOffice:    boxOffice,
		TotalSeasons: totalSeasons,
	}
}
