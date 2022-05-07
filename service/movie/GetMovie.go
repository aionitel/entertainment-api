package service

import (
	"github.com/gin-gonic/gin"
)

// get movie by name in query param

func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieTitle := c.Query("title")

		println(movieTitle)
	}
}
