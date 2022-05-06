package routes

import (
	"github.com/gin-gonic/gin"
	services "github.com/oranges0da/go-server/services/movie"
)

func Movie() *gin.RouterGroup {
	r := gin.New()

	routes := r.Group("/")
	{
		routes.GET("/:name", services.GetMovies())
	}

	return routes
}
