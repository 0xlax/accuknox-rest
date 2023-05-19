package main

import (
	"log"

	"github.com/0xlax/accuknox-rest/cmd/routes"
	"github.com/0xlax/accuknox-rest/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
