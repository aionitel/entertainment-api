package service

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oranges0da/entertainment-api/model"
	"github.com/oranges0da/entertainment-api/service/utils"
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

	if err != nil {
		panic(err)
	}

	defer res.Body.Close() // close the body after reading

	data := utils.UnpackMovie(res) // utils function that unmarshalls the res

	return data
}

// get movie by imdb id in query param
func GetMovieByID(client *http.Client, id string) {

}

// main function that sends requests using HttpClient
func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieTitle := c.Query("title")

		if movieTitle != "" {
			movieData := GetMovieByTitle(utils.HttpClient(), movieTitle)

			c.JSON(200, gin.H{
				"data": movieData,
			})
		}
	}
}
