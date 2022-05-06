package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/oranges0da/go-server/service/movie"
)

func RunServer(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func main() {
	router := gin.New()

	router.Group("/movie")
	{
		router.GET("/:name", service.GetMovie())
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	RunServer(server)
}
