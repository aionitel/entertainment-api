package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oranges0da/entertainment-api/handler"
)

func Config() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	log.Printf("Running with %d CPUs\n", nuCPU)
}

// RunServer start server and quit if error occurs or user quits
func Run(server *http.Server) {
	log.Printf("Running on port: %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func main() {
	godotenv.Load(".env")
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))
	handler.NewHandler(&handler.Config{
		R: router,
	})

	// get port from env, if no port value is found, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	Run(server)
}
