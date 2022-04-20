package main

import (
	"github.com/Maritho/go-fibers-learning/databases"
	"github.com/Maritho/go-fibers-learning/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	databases.ConnectDB()

	router.SetupRoutes(app)

	// Listen on PORT 300
	app.Listen(":8000")
}
