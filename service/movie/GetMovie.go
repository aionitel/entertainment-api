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
	key := os.Getenv("MOVIE_KEY")
	println("Get movie hit")

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

	data := utils.Unpack(res, model.Movie{}) // utils function that unmarshalls the res

	return data
}

// get movie by imdb id in query param
func GetMovieByID(client *http.Client, id string) model.Movie {
	key := os.Getenv("MOVIE_KEY")

	var url string = "http://www.omdbapi.com/?i=" + id + "&apikey=" + key
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

	data := utils.Unpack(res, model.Movie{}) // utils function that unmarshalls the res

	return data
}

// main function that sends requests using HttpClient
func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get query params for either title or imdb id
		movieTitle := c.Query("q")
		imdb := c.Query("imdb")

		if movieTitle != "" { // send request by title
			movieData := GetMovieByTitle(utils.HttpClient(), movieTitle)

			c.JSON(200, gin.H{
				"data":   movieData,
				"errors": []string{},
			})
		} else if imdb != "" { // send request by imdb id
			movieData := GetMovieByID(utils.HttpClient(), imdb)

			c.JSON(200, gin.H{
				"data":   movieData,
				"errors": []string{},
			})
		} else {
			c.JSON(404, gin.H{
				"data":   nil,
				"errors": []string{"No movie title or imdb id provided"},
			})
		}
	}
}
