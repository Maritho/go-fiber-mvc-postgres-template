package router

import (
	noteRoutes "github.com/Maritho/go-fiber-mvc-postgres-template/app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Setup the Node Routes
	noteRoutes.SetupRoutes(api)
}
