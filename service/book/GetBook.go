package service

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oranges0da/entertainment-api/model"
	"github.com/oranges0da/entertainment-api/service/utils"
)

func GetMovieByTitle(client *http.Client, title string) model.Book {
	key := os.Getenv("BOOK_KEY") // get key to book api from env
	var url string = "https://www.googleapis.com/books/v1/volumes?q=" + title + "&apikey=" + key
	log.Print(url)

	req, err := http.NewRequest("GET", url, nil) // create request, not send it

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req) // send request

	if err != nil {
		panic(err)
	}

	defer res.Body.Close() // close the body after reading

	data := utils.UnpackBook(res)

	return data
}

func GetBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookTitle := c.Query("title") // get title from url

		data := GetMovieByTitle(utils.HttpClient(), bookTitle) // get movie by title

		c.JSON(http.StatusOK, gin.H{
			"data":   data,
			"errors": []string{},
		})
	}
}
