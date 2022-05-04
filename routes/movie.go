package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oranges0da/go-server/services"
)

func Movie() *gin.RouterGroup {
	r := gin.New()

	routes := r.Group("/movie")
	{
		routes.GET("/", services.GetHome())
	}

	return routes
}
