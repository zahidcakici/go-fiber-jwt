package controller

import (
	"go-fiber-jwt-example/database"
	"go-fiber-jwt-example/model"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	var books []model.Book
	database.DB.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book model.Book
	database.DB.Find(&book, id)
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	var book bookRequest
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	bookModel := model.Book{
		Title:  book.Title,
		Author: book.Author,
	}
	resp := database.DB.Create(&bookModel)
	if resp.Error != nil {
		return c.Status(400).SendString(resp.Error.Error())
	}
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book bookRequest
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	var oldBook model.Book
	database.DB.Find(&oldBook, id)
	if oldBook.Title == "" {
		return c.Status(404).SendString("No book found with ID")
	}
	database.DB.Model(&oldBook).Updates(book)
	return c.JSON(oldBook)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book model.Book
	database.DB.Find(&book, id)
	if book.Title == "" {
		return c.Status(404).SendString("No book found with ID")
	}
	database.DB.Delete(&book)
	return c.SendString("Book successfully deleted")
}

type bookRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}
