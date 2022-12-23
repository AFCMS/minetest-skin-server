package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"

	"github.com/gofiber/fiber/v2"
)

// Return the skin file
func SkinFull(c *fiber.Ctx) error {
	var skin models.Skin

	if err := database.DB.Find(&skin, "").Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	c.Status(fiber.StatusOK)
	c.Set("Content-Type", "image/png")

	return c.Send(skin.Data)
}
