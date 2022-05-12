package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbum(client *http.Client, title string) gin.HandlerFunc {
	var url string = "https://api.deezer.com/search?q=\"" + title + "\""

	log.Print(url)

	return func(c *gin.Context) {

	}
}
