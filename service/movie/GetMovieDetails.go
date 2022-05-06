package service

import "github.com/gin-gonic/gin"

func GetMovieDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieName := c.Param("name")

		// query api for movie details

		movie := model.NewMovie(
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
	}
}