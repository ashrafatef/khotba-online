package masjeds

import (
	"khotba-online/internal/database"
	"khotba-online/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupMasjedRoutes(app fiber.Router) {
	masjed := app.Group("/masjeds")
	masjed.Use(middlewares.AuthMiddleware())
	repo := NewMasjedRepository(database.New())
	controller := NewMasjedController(repo)

	// Protected routes
	masjed.Post("/", controller.Create)
}
