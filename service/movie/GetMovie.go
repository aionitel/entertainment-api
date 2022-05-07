package service

import (
	"github.com/gin-gonic/gin"
	"github.com/oranges0da/go-server/model"
)

// get movie by name in query param

func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		// movieName := c.Param("name")

		// query api for movie details

		movie := model.NewMovie(
			1,
			"2001",
			"Two imprisoned",
			"Stanley Kubrick",
			1968,
			8.3,
		)

		c.JSON(200, gin.H{
			"movies": movie,
		})
	}
}
