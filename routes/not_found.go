package routes

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotFound)
}
