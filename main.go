package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oranges0da/go-server/routes"
)

func main() {
	app := echo.New()

	app.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "Hello World")
	})

	app.GET("/home", routes.Home)

	app.Logger.Fatal(app.Start(":4000"))
}
