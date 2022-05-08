package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// get movie details by name in query param
func GetMovieByTitle(client *http.Client, title string) {
	key := os.Getenv("API_KEY")

	var url string = "http://www.omdbapi.com/?t=" + title + "&apikey=" + key

	log.Print(url)

	req, err := http.NewRequest("GET", url, nil) // create request, not send it

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req) // send request

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	log.Print(string(body))
}

// get movie by imdb id in query param
func GetMovieByID(id string) {

}

// main function that sends requests using HttpClient
func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieTitle := c.Query("title")

		GetMovieByTitle(http.DefaultClient, movieTitle)

		c.JSON(200, gin.H{
			"data": movieTitle,
		})
	}
}
