package main

import (
	"log"

	"github.com/0xlax/accuknox-rest/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Note taking Application")
	})

	log.Fatal(app.Listen(":3000"))
}
