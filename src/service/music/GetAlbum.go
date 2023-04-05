package service

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oranges0da/entertainment-api/src/model"
	"github.com/oranges0da/entertainment-api/src/service/utils"
)

func GetAlbumByTitle(client *http.Client, title string) model.Music {
	var url string = "https://api.deezer.com/search?q=album:\"" + title + "\""
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

func GetAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		albumQuery := c.Query("q")
		albumTitle := strings.ReplaceAll(albumQuery, " ", "+")

		data := GetSongByTitle(utils.HttpClient(), albumTitle)

		c.JSON(http.StatusOK, data)
	}
}
