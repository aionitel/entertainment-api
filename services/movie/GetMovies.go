package services

import "github.com/gin-gonic/gin"

func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"movies": []string{"2001: Space Odyssey", "Back to the Future"},
		})
	}
}
