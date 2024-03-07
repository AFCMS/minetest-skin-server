package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"
	"minetest-skin-server/types"

	"github.com/gofiber/fiber/v2"
)

func SkinList(c *fiber.Ctx) error {
	// Parse query
	queryR := types.QuerySkinList{}
	if err := c.QueryParser(&queryR); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Bad request")
	}

	// Convert values to handle GORM API
	count := int(queryR.Count)

	if queryR.Count == 0 {
		count = -1
	}

	page := int(queryR.Page)

	if queryR.Page == 0 {
		page = -1
	}

	// Query database
	var result []models.Skin

	if err := database.DB.Limit(count).Offset(page * count).Find(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Cannot interact with database")
	}

	return c.JSON(result)
}
