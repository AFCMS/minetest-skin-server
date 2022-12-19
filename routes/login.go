package routes

import (
	"minetest-skin-server/types"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	input := types.InputLogin{}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	return nil
}
