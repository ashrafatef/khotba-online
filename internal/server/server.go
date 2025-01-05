package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"khotba-online/internal/database"
	"khotba-online/pkg/errors"
)

type FiberServer struct {
	*fiber.App

	db *gorm.DB
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "khotba-online",
			AppName:      "khotba-online",
			ErrorHandler: func(c *fiber.Ctx, err error) error {

				if validationErr, ok := err.(*errors.ValidationError); ok {
					return c.Status(validationErr.StatusCode).JSON(fiber.Map{
						"message": validationErr.Message,
						"fields":  validationErr.Fields,
					})
				}
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			},
		}),

		db: database.New(),
	}

	return server
}
