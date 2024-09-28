package main

import (
	"go-fiber-jwt-example/database"
	"go-fiber-jwt-example/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	router.SetupRoutes(app)

	log.Fatal(app.Listen("localhost:3000"))
}
