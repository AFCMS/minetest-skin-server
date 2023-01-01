package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Return the skin file
func UsersID(c *fiber.Ctx) error {
	var a models.Account

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	a, err = database.AccountFromID(uint(id))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(a)
}
