package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"

	"github.com/gofiber/fiber/v2"
)

// TODO: cache result
func SkinRecent(c *fiber.Ctx) error {
	var results []models.Skin
	if err := database.DB.Find(&results).Order("created_at DESC").Limit(10).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(results)
}
