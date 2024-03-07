package routes

import (
	"minetest-skin-server/database"

	"github.com/gofiber/fiber/v2"
)

// TODO: include user and skin count

func Info(c *fiber.Ctx) error {
	accountCount, err := database.AccountCount()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	skinCount, err := database.SkinCount()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"version":       "1.0",
		"account_count": accountCount,
		"skin_count":    skinCount,
	})
}
