package services

import (
	"github.com/gin-gonic/gin"
	"github.com/oranges0da/go-server/models"
)

func GetMovie(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieName := c.Param("name")

		// query api for movie details

		movie := models.NewMovie({})

		c.JSON(200, gin.H{
			"movies": movie,
		})
	}
}
