package handler

import (
	"github.com/0xlax/accuknox-rest/database"
	"github.com/0xlax/accuknox-rest/models"
	"github.com/gofiber/fiber/v2"
)

// Home is the handler function for the home route ("/").
// It sends a string response "Note taking Application".
func ListNotes(c *fiber.Ctx) error {
	notes := []models.Word{}
	database.DB.Db.Find(&notes)
	// return c.SendString("Note taking Application")
	return c.Status(200).JSON(notes)
}

func CreateNote(c *fiber.Ctx) error {
	note := new(models.Word)
	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&note)

	return c.Status(200).JSON(note)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "400 BAD REQUEST",
		})
	}

	database.DB.Db.Create(&user)
	return c.Status(fiber.StatusCreated).JSON(user)
}
