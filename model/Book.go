package model

type Book struct {
	Title          string   `json:"title"`
	Subtitle       string   `json:"subtitle"`
	Authors        []string `json:"authors"`
	Publisher      string   `json:"publisher"`
	PublishedDate  string   `json:"publishedDate"`
	Desc           string   `json:"description"`
	PageCount      int      `json:"pageCount"`
	PrintType      string   `json:"printType"`
	Categories     []string `json:"categories"`
	AverageRating  float32  `json:"averageRating"`
	MaturityRating string   `json:"maturityRating"`
	Cover          string   `json:"imageLinks"`
	Language       string   `json:"language"`
	PublicDomain   bool     `json:"publicDomain"`
}

func NewBook(
	title string,
	subtitle string,
	authors []string,
	publisher string,
	publishedDate string,
	desc string,
	pageCount int,
	printType string,
	categories []string,
	averageRating float32,
	maturityRating string,
	cover string,
	language string,
	publicDomain bool,
) *Book {
	return &Book{
		Title:          title,
		Subtitle:       subtitle,
		Authors:        authors,
		Publisher:      publisher,
		PublishedDate:  publishedDate,
		Desc:           desc,
		PageCount:      pageCount,
		PrintType:      printType,
		Categories:     categories,
		AverageRating:  averageRating,
		MaturityRating: maturityRating,
		Cover:          cover,
		Language:       language,
		PublicDomain:   publicDomain,
	}
}
