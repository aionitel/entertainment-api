package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oranges0da/entertainment-api/model"
)

// get movie by name in query param
func GetMovieByTitle(client *http.Client, title string) *http.Response {
	key := os.Getenv("API_KEY")

	var url string = "http://www.omdbapi.com/?apikey=[" + key + "]&t=" + title

	req, err := http.NewRequest("GET", url, nil) // create request, not send it

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req) // send request

	if err != nil {
		panic(err)
	}

	defer res.Body.Close() // close body after request is done

	body, err := ioutil.ReadAll(res.Body) // read body

	data := json.Unmarshal(body, &model.Movie)

	return res
}

// get movie by imdb id in query param
func GetMovieByID(id string) {

}

// main function that sends requests using HttpClient
func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieTitle := c.Query("title")

		res := GetMovieByTitle(http.DefaultClient, movieTitle)

		c.JSON(200, gin.H{
			"data": string(res),
		})
	}
}
