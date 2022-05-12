package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oranges0da/entertainment-api/model"
	"github.com/oranges0da/entertainment-api/service/utils"
)

func GetTrackByTitle(client *http.Client, title string) model.Music {
	var url string = "https://api.deezer.com/track/search?q=title:\"" + title + "\""
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

	data := utils.UnpackTrack(res) // utils function that unmarshalls the res

	return data
}

func GetTrack() gin.HandlerFunc {
	return func(c *gin.Context) {
		trackTitle := c.Param("q")

		data := GetTrackByTitle(utils.HttpClient(), trackTitle)

		c.JSON(http.StatusOK, data)
	}
}
