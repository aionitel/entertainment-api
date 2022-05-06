package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	server.ListenAndServe()
}
