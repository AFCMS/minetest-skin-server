package routes

import (
	"luanti-skin-server/database"
	"luanti-skin-server/models"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// SkinHead Return the skin file
func SkinHead(c fiber.Ctx) error {
	var skin models.Skin

	id, err := uuid.Parse(c.Params("uuid"))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	skin, err = database.SkinFromUUID(id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	c.Status(fiber.StatusOK)
	c.Set("Content-Type", "image/png")

	return c.Send(skin.DataHead)
}
