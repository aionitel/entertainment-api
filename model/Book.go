package model

type VolumeInfo struct {
	Title          string      `json:"title"`
	Authors        []string    `json:"authors"`
	Publisher      string      `json:"publisher"`
	PublishedDate  string      `json:"publishedDate"`
	Desc           string      `json:"description"`
	PageCount      int         `json:"pageCount"`
	PrintType      string      `json:"printType"`
	AverageRating  float32     `json:"averageRating"`
	MaturityRating string      `json:"maturityRating"`
	ImageLinks     CoverImages `json:"imageLinks"`
	Language       string      `json:"language"`
}

type CoverImages struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

type Item struct {
	Kind       string     `json:"kind"`
	ID         string     `json:"id"`
	Etag       string     `json:"etag"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
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
