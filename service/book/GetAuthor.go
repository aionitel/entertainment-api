package service

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oranges0da/entertainment-api/model"
	"github.com/oranges0da/entertainment-api/service/utils"
)

func GetAuthorByName(client *http.Client, name string) model.Book {
	key := os.Getenv("BOOK_KEY") // get key to book api from env
	var url string = "https://www.googleapis.com/books/v1/volumes?q=" + name + "&key=" + key + "&inauthor"
	log.Print(url)

	req, err := http.NewRequest("GET", url, nil) // create request, not send it

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req) // send request

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close() // close the body after reading

	data := utils.Unpack(res, model.Book{})

	return data
}

func GetAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorQuery := c.Query("q")                         // get author from url
		author := strings.ReplaceAll(authorQuery, " ", "+") // replace space with + for annoying query

		if author == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Author's name is required.",
			})
		}
		data := GetAuthorByName(utils.HttpClient(), author)

		c.JSON(http.StatusOK, gin.H{
			"data":   data,
			"errors": []string{},
		})
	}
}
