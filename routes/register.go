package routes

import (
	"minetest-skin-server/types"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	input := types.InputRegister{}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on register request", "data": err})
	}
	return nil
}
