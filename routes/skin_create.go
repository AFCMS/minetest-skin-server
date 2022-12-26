package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"
	"minetest-skin-server/types"
	"minetest-skin-server/utils"

	"github.com/gofiber/fiber/v2"
)

// Handle Skin creation
//
// Use a multipart request
func SkinCreate(c *fiber.Ctx) error {
	input := types.InputSkinCreate{}

	// Get the text fields
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on create request", "data": err})
	}

	// Get file part
	var b []byte
	var err error
	if b, err = utils.LoadFormFile(c, "data"); err != nil {
		return err
	}

	// TODO: test decoding as PNG
	// TODO: validate file dimensions
	// TODO: run optipng

	input.Data = b

	//log.Println(input.Data)

	var l = models.Skin{
		Description: input.Description,
		Public:      input.Public,
		Data:        input.Data,
	}

	if err := database.DB.Create(&l).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Cannot interact with database", "data": err})
	}

	return c.SendStatus(fiber.StatusNotImplemented)
}
