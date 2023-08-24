package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ionutinit/riddles-api-go/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	log.Fatal(app.Listen(":3000"))
}
