package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"

	"github.com/gofiber/fiber/v2"
)

func SkinList(c *fiber.Ctx) error {

	//count, _ := c.ParamsInt("count", 40)

	//search := c.Params("search", "")

	var result []models.Skin

	if err := database.DB.Find(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Cannot interact with database")
	}

	return c.JSON(result)

	// return c.SendStatus(fiber.StatusNotImplemented)
}
