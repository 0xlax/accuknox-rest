package main

import (
	"github.com/0xlax/accuknox-rest/handler"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", handler.Home)

}
