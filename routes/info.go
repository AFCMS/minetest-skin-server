package routes

import (
	"github.com/gofiber/fiber/v3"

	"minetest-skin-server/database"
	"minetest-skin-server/utils"
)

// TODO: include user count

func Info(c fiber.Ctx) error {
	accountCount, err := database.AccountCount()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	skinCount, err := database.SkinCount()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return supported OAuth providers
	// Used by the frontend to determine which OAuth buttons to display
	var supportedOAuthProviders []string

	if utils.ConfigOAuthContentDB {
		supportedOAuthProviders = append(supportedOAuthProviders, "contentdb")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"version":                   "1.0",
		"account_count":             accountCount,
		"skin_count":                skinCount,
		"supported_oauth_providers": supportedOAuthProviders,
	})
}
