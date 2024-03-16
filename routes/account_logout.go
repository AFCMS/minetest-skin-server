package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func AccountLogout(c *fiber.Ctx) error {
	err := c.Locals("session").(*session.Session).Destroy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Cannot logout"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logged out"})
}
