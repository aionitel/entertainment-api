package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oranges0da/go-server/handler"
)

func RunServer(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func main() {
	router := gin.New()

	handler.NewHandler(&handler.Config{
		R: router,
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	RunServer(server)
}
