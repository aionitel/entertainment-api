package service

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oranges0da/entertainment-api/model"
	"github.com/oranges0da/entertainment-api/service/utils"
)

func GetTVByTitle(client *http.Client, title string) model.TV {
	key := os.Getenv("TV_KEY")

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

	data := utils.UnpackTV(res) // utils function that unmarshalls the res

	return data
}

func GetTVByID(client *http.Client, id string) model.TV {
	key := os.Getenv("API_KEY")

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

	data := utils.UnpackTV(res) // utils function that unmarshalls the res

	return data
}

func GetTV() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get query params for either title or imdb id
		tvTitle := c.Query("title")
		imdb := c.Query("imdb")

		if tvTitle != "" { // send request by title
			tvData := GetTVByTitle(utils.HttpClient(), tvTitle)

			c.JSON(200, gin.H{
				"data":   tvData,
				"errors": []string{},
			})
		} else if imdb != "" { // send request by imdb id
			tvData := GetTVByID(utils.HttpClient(), imdb)

			c.JSON(200, gin.H{
				"data":   tvData,
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
