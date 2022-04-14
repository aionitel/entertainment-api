package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oranges0da/go-server/routes"
)

func main() {
	router := gin.New()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", routes.Home)

	router.Run(":4000")
}
