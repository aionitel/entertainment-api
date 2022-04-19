package routes

import "github.com/gin-gonic/gin"

func Contact() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "contact page",
		})
	}
}
