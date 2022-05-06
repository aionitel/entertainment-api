package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oranges0da/go-server/routes"
)

func RunServer(server *http.Server) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to initialize server.")
	}

}

func main() {
	app := gin.New()

	server := &http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	app.Group("/movie")
	{
		app.GET("/:name", routes.MovieRoutes())
	}
}
