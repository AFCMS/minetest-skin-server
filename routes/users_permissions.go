package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"
	"minetest-skin-server/types"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// UsersPermissions Set user permissions
func UsersPermissions(c *fiber.Ctx) error {
	var a models.Account

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	a, err = database.AccountFromID(uint(id))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	input := types.InputUsersPermissions{}

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	database.AccountSetPermission(&a, input.Level)

	return c.SendStatus(fiber.StatusOK)
}
