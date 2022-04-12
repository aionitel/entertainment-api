package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oranges0da/go-server/routes"
)

func main() {
	app := fiber.New(fiber.Config{ // root app with config
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		GETOnly:       true, // for now
	})

	app.Get("/contact", routes.Contact)

	app.Listen(":4000")
}
