package service

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oranges0da/entertainment-api/model"
	"github.com/oranges0da/entertainment-api/service/utils"
)

func GetArtistByName(client *http.Client, name string) model.Music {
	var url string = "https://api.deezer.com/search?q=artist:\"" + name + "\""
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

func GetArtist() gin.HandlerFunc {
	return func(c *gin.Context) {
		artistQuery := c.Query("q")
		artistName := strings.ReplaceAll(artistQuery, " ", "+")

		data := GetSongByTitle(utils.HttpClient(), artistName)

		c.JSON(http.StatusOK, data)
	}
}
