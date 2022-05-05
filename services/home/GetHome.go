package services

import "github.com/gin-gonic/gin"

func GetHome() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "hello",
		})
	}
}
