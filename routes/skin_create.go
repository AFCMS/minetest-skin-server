package routes

import (
	"log"
	"minetest-skin-server/types"

	"github.com/gofiber/fiber/v2"
)

// Handle Skin creation
//
// Use a multipart request
func SkinCreate(c *fiber.Ctx) error {
	input := types.InputSkinCreate{}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on create request", "data": err})
	}

	log.Println(input.Data)

	return c.SendStatus(fiber.StatusNotImplemented)
}
