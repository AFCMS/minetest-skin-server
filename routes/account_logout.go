package routes

import "github.com/gofiber/fiber/v2"

func AccountLogout(c *fiber.Ctx) error {
	c.ClearCookie()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logged out"})
}
