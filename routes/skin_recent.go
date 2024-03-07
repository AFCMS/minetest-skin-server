package routes

import (
	"minetest-skin-server/database"

	"github.com/gofiber/fiber/v2"
)

// SkinRecent TODO: cache result
func SkinRecent(c *fiber.Ctx) error {
	results, err := database.SkinRecent(10)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(results)
}
