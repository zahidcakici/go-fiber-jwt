package router

import (
	"go-fiber-jwt-example/controller"
	"go-fiber-jwt-example/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	// Book
	book := api.Group("/books")
	book.Get("/", controller.GetBooks)
	book.Get("/:id", controller.GetBook)

	book.Use(middleware.JWTProtected)
	book.Post("/", controller.CreateBook)
	book.Patch("/:id", controller.UpdateBook)
	book.Delete("/:id", controller.DeleteBook)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/register", controller.Register)
}
