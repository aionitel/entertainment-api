package service

import (
	"github.com/gin-gonic/gin"
)

// get movie by name in query param

func GetMovieByTitle(title string) {

}

func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieTitle := c.Query("title")
		imdbID := c.Query("imdb")

	}
}
