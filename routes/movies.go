package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/oranges0da/go-server/services/movie"
)

func GetMovies() *gin.RouterGroup {
	r := gin.New()

	routes := r.Group("/movies")
	{
		routes.GET("/", services.GetMovies())
	}

	return routes
}
