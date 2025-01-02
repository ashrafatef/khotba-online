package server

import (
	"github.com/gofiber/fiber/v2"

	"khotba-online/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "khotba-online",
			AppName:      "khotba-online",
		}),

		db: database.New(),
	}

	return server
}
