package emam

import (
	"khotba-online/internal/database"

	"github.com/gofiber/fiber/v2"
)

func AddEmamRouters(app *fiber.App) {
	emam := app.Group("/emams")
	repo := NewEmamRepository(database.New())
	controller := NewEmamController(repo)

	// routers
	emam.Post("/signup", controller.SignUp)
	emam.Post("/login", controller.Login)
}
