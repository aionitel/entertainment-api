package model

type VolumeInfo struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}

type CoverImages struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

type Item struct {
	VolumeInfo     VolumeInfo  `json:"volumeInfo"`
	Publisher      string      `json:"publisher"`
	PublishedDate  string      `json:"publishedDate"`
	PageCount      string      `json:"pageCount"`
	PrintType      string      `json:"printType"`
	MaturityRating string      `json:"maturityRating"`
	ImageLinks     CoverImages `json:"imageLinks"`
	Language       string      `json:"language"`
}

type Book struct {
	Kind       string `json:"kind"`
	TotalItems int    `json:"totalItems"`
	Items      []Item `json:"items"`
}

func NewBook(
	kind string,
	totalItems int,
	items []Item,
) *Book {
	return &Book{
		Kind:       kind,
		TotalItems: totalItems,
		Items:      items,
	}
}
