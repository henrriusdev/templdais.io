package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		AppName:       "TemplDais",
		CaseSensitive: true,
	})

	app.Use(logger.New())
	app.Use(ErrorHandler())

	// serve static folder with fiber
	app.Static("/static", "usr/local/bin/static", fiber.Static{
		Browse: true,
	})

	app.Get("/", home)
	app.Get("/components/:component", components)

	log.Fatal(app.Listen(":8080"))
}
