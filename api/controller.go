package api

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/doemoor/moci/internal/database"
)

func GetAllControllers(c *fiber.Ctx) error {

	return c.SendString("Hello, GetAllControllers 👋!")
}

func GetControllerById(c *fiber.Ctx) error {
	return c.SendString("Hello, GetControllerById 👋!")
}