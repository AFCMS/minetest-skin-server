package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"

	"luanti-skin-server/database"
	"luanti-skin-server/models"
)

func SkinApprove(c fiber.Ctx) error {
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
