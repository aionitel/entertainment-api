package handler

import (
	"github.com/gin-gonic/gin"
	services "github.com/oranges0da/go-server/services/movie"
)

type Handler struct{}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}

	g := c.R.Group("/api")
	{
		g.GET("/movies/:name", services.GetMovie())
	}
}
