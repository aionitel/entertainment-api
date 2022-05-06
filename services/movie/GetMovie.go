package services

import (
	"github.com/gin-gonic/gin"
	"github.com/oranges0da/go-server/models"
)

func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieName := c.Param("name")

		// query api for movie details

		movie := models.NewMovie(
			1,
			movieName,
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
