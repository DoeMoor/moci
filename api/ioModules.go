package api

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/doemoor/moci/internal/database"
)

func GetAllIoModules(c *fiber.Ctx) error {
	return c.SendString("Hello, GetAllIoModules 👋!")
}

func GetIoModuleById(c *fiber.Ctx) error {
	return c.SendString("Hello, GetIoModuleById 👋!")
}