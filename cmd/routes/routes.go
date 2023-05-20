// routes/routes.go
package routes

import (
	"github.com/0xlax/accuknox-rest/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes is responsible for configuring the routes of the application.
// It accepts a fiber.App instance and sets up the necessary routes and their corresponding handler functions.
func SetupRoutes(app *fiber.App) {

	app.Post("/signup", handler.CreateUser)
	app.Post("/login", handler.Login)
	app.Post("/notes", handler.CreateNoteToUser)
	app.Get("/notes", handler.ListNotes)
	app.Delete("/notes", handler.DeleteNote)

}
