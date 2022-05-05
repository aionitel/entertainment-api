package main

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func BundleRoutes() routes {
	r := routes{
		router: gin.New(),
	}

	movie := r.router.Group("/movie")
}

func main() {
	app := gin.New()

}
