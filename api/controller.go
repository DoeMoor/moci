package api

import "github.com/gofiber/fiber/v3"

func GetAllControllers(c fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}