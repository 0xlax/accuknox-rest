package handler

import "github.com/gofiber/fiber/v2"

// Home is the handler function for the home route ("/").
// It sends a string response "Note taking Application".
func Home(c *fiber.Ctx) error {
	return c.SendString("Note taking Application")
}
