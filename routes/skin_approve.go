package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SkinApprove(c *fiber.Ctx) error {
	var skin models.Skin

	id, err := uuid.Parse(c.Params("uuid"))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	skin, err = database.SkinFromUUID(id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	err = database.SkinApproval(&skin, true)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON("Sucess")
}
