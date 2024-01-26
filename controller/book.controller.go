package controller

import (
	"bookapi.com/database"
	"bookapi.com/model"
	"github.com/gofiber/fiber/v2"
)

// ShowAccount godoc
// @Summary Show a account
// @Description get all book
// @Accept  json
// @Produce  json
// @Router /api/book [get]
func GetAllBook(c *fiber.Ctx) error {
	db := database.DB.Db
	var books []model.Book
	db.Find(&books)
	if len(books) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "success", "message": "Books is empty", "data": books})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Books Found", "data": books})
}

func Create(c *fiber.Ctx) error {
	db := database.DB.Db
	book := new(model.Book)
	err := c.BodyParser(book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&book).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	// Return the created book
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": book})
}
