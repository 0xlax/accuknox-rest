package handler

import (
	"github.com/0xlax/accuknox-rest/database"
	"github.com/0xlax/accuknox-rest/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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

func Login(c *fiber.Ctx) error {
	loginData := new(models.LoginUser)
	if err := c.BodyParser(loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "400 BAD REQUEST",
		})
	}

	// Find the user by email
	user := database.DB.FindUserByEmail(loginData.Email)
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// Check if the password matches
	if user.Password != loginData.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// Generate a unique session ID (sid) for the user
	sid := uuid.New().String()

	// Update the user's sessionID
	user.SessionID = sid
	database.DB.Db.Save(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logged in successfully",
		"sid":     sid,
	})
}

func CreateNoteToUser(c *fiber.Ctx) error {
	requestBody := new(struct {
		SID   string `json:"sid"`
		Notes string `json:"notes"`
	})
	if err := c.BodyParser(requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Authenticate the user based on the session ID
	user := database.DB.FindUserBySessionID(requestBody.SID)
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not authenticated",
		})
	}

	// Create a new Word object with the provided notes and user's ID
	note := models.Word{
		Notes:  requestBody.Notes,
		UserID: user.ID,
	}

	// Save the note in the database
	database.DB.Db.Create(&note)

	// Append the note to the user's Notes array
	user.Notes = append(user.Notes, note)

	// Save the updated user in the database
	database.DB.Db.Save(&user)

	// Get the ID of the newly created note
	newNoteID := note.ID

	// Return only the ID of the newly created note in the response
	response := fiber.Map{
		"note_id": newNoteID,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func ListNotes(c *fiber.Ctx) error {
	requestBody := new(struct {
		SID string `json:"sid"`
	})
	if err := c.BodyParser(requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Authenticate the user based on the session ID
	user := database.DB.FindUserBySessionID(requestBody.SID)
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not authenticated",
		})
	}

	// Retrieve the notes associated with the user
	notes := []models.Word{}
	database.DB.Db.Where("user_id = ?", user.ID).Find(&notes)

	// Create a response array with notes and their IDs
	response := make([]map[string]interface{}, len(notes))
	for i, note := range notes {
		response[i] = map[string]interface{}{
			"id":   note.ID,
			"note": note.Notes,
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"notes": response,
	})
}
