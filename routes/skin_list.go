package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"
	"minetest-skin-server/types"

	"github.com/gofiber/fiber/v3"
)

func SkinList(c fiber.Ctx) error {
	// Parse query
	queryR := new(types.QuerySkinList)
	if err := c.Bind().Query(queryR); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Bad request")
	}

	// Convert values to handle GORM API
	count := int(queryR.Count)

	if queryR.Count == 0 {
		count = 10
	}

	page := int(queryR.Page)

	if queryR.Page == 0 {
		page = 0
	}

	// Query database
	var result []models.Skin

	if err := database.DB.Limit(count).Offset(page * count).Find(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Cannot interact with database")
	}

	return c.JSON(result)
}
