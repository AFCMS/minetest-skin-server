package routes

import (
	"minetest-skin-server/models"

	"github.com/gofiber/fiber/v2"
)

func AccountUser(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Account)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":               user.ID,
		"name":             user.Name,
		"email":            user.Email,
		"permission_level": user.PermissionLevel,
	})
}
