package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oranges0da/entertainment-api/model"
)

// get movie details by name in query param
func GetMovieByTitle(client *http.Client, title string) model.Movie {
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

	var data model.Movie

	err = json.Unmarshal(body, &data)

	return data
}

// get movie by imdb id in query param
func GetMovieByID(id string) {

}

// main function that sends requests using HttpClient
func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieTitle := c.Query("title")

		data := GetMovieByTitle(http.DefaultClient, movieTitle)

		c.JSON(200, gin.H{
			"data": data,
		})
	}
}
