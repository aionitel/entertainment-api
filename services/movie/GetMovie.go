package services

import "github.com/gin-gonic/gin"

func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"movie": name,
		})
	}
}
