package routes

import (
	"minetest-skin-server/database"

	"github.com/gofiber/fiber/v2"
)

// TODO: include user and skin count

func Info(c *fiber.Ctx) error {
	account_count, err := database.AccountCount()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	skin_count, err := database.SkinCount()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"version":       "1.0",
		"account_count": account_count,
		"skin_count":    skin_count,
	})
}
