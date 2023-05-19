// routes/routes.go
package routes

import (
	"github.com/0xlax/accuknox-rest/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes is responsible for configuring the routes of the application.
// It accepts a fiber.App instance and sets up the necessary routes and their corresponding handler functions.
func SetupRoutes(app *fiber.App) {

	app.Get("/", handler.ListNotes)
	app.Post("/signup", handler.CreateUser)
	app.Post("/notes", handler.CreateNote)
	app.Post("/login", handler.Login)

}

// POST 	/signup
// POST 	/login
// GET  	/notes
// POST 	/notes
// DELETE   /notes
