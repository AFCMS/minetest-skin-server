package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Info(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"version": "1.0"})
}
