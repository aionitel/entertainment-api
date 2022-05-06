package handler

import (
	"github.com/gin-gonic/gin"
	service "github.com/oranges0da/go-server/service/movie"
)

type Handler struct{} // handling requests and allocating services towards them

type Config struct {
	R *gin.Engine // for configurating root router
}

func NewHandler(c *Config) {
	//h := &Handler{}

	g := c.R.Group("/api") // root endpoint

	g.GET("/movies/:name", service.GetMovie()) // sample endpoint
}
