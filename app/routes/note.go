package routes

import (
	noteHandler "github.com/Maritho/go-fiber-mvc-postgres-template/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	note := router.Group("/note")
	// Create a Note
	note.Post("/", noteHandler.CreateNotes)
	// Read all Notes
	note.Get("/", noteHandler.GetNotes)
	// Read one Note
	note.Get("/:noteId", noteHandler.GetNote)
	// Update one Note
	note.Put("/:noteId", noteHandler.UpdateNote)
	// Delete one Note
	note.Delete("/:noteId", noteHandler.DeleteNote)
}
