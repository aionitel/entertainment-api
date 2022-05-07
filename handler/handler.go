package handler

import (
	"github.com/gin-gonic/gin"
	service "github.com/oranges0da/entertainment-api/service/movie"
)

type Config struct {
	R *gin.Engine // for configurating root router
}

func NewHandler(c *Config) {
	g := c.R.Group("/movie") // movie endpoint
	{
		g.GET("/:name", service.GetMovie())
	}
}
