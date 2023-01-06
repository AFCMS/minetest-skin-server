package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"
	"minetest-skin-server/types"

	"github.com/gofiber/fiber/v2"
)

func SkinList(c *fiber.Ctx) error {
	// Parse query
	query_r := types.QuerySkinList{}
	if err := c.QueryParser(&query_r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Bad request")
	}

	// Convert values to handle GORM API
	count := int(query_r.Count)

	if query_r.Count == 0 {
		count = -1
	}

	page := int(query_r.Page)

	if query_r.Page == 0 {
		page = -1
	}

	// Query database
	var result []models.Skin

	if err := database.DB.Limit(count).Offset(page * count).Find(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Cannot interact with database")
	}

	return c.JSON(result)
}
