package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ionutinit/riddles-api-go/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
