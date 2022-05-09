package handler

import (
	"github.com/gin-gonic/gin"
	movie "github.com/oranges0da/entertainment-api/service/movie"
	tv "github.com/oranges0da/entertainment-api/service/tv"
)

type Config struct {
	R *gin.Engine // for configurating root router
}

func NewHandler(c *Config) {
	movieGroup := c.R.Group("/movie") // movie endpoint
	{
		movieGroup.GET("/", movie.GetMovie())
	}

	TVGroup := c.R.Group("/tv") // tv endpoint
	{
		TVGroup.GET("/", tv.GetTV())
	}
}
