package routes

import "github.com/gin-gonic/gin"

func Contact(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
