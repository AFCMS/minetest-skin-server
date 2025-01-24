package routes

import (
	"luanti-skin-server/database"
	"luanti-skin-server/models"
	"luanti-skin-server/types"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// UsersPermissions Set user permissions
func UsersPermissions(c fiber.Ctx) error {
	var a models.Account

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	a, err = database.AccountFromID(uint(id))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	input := new(types.InputUsersPermissions)

	if err := c.Bind().JSON(input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = database.AccountSetPermission(&a, input.Level)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
