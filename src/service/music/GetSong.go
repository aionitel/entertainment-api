package service

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oranges0da/entertainment-api/model"
	"github.com/oranges0da/entertainment-api/service/utils"
)

func GetSongByTitle(client *http.Client, title string) model.Music {
	var url string = "https://api.deezer.com/search?q=track:\"" + title + "\""
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

	data := utils.Unpack(res, model.Music{}) // utils function that unmarshalls the res

	return data
}

func GetSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		trackQuery := c.Query("q")
		trackTitle := strings.ReplaceAll(trackQuery, " ", "+")

		data := GetSongByTitle(utils.HttpClient(), trackTitle)

		c.JSON(http.StatusOK, data)
	}
}
