package routes

import (
	"bytes"
	"image/png"
	"minetest-skin-server/database"
	"minetest-skin-server/models"
	"minetest-skin-server/types"
	"minetest-skin-server/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Handle Skin creation
//
// Use a multipart request
func SkinCreate(c *fiber.Ctx) error {
	// Get User
	user := c.Locals("user").(models.Account)

	input := types.InputSkinCreate{}

	// Get the text fields
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on create request", "data": err})
	}

	// Get file part
	var skin_b []byte
	var err error
	if skin_b, err = utils.LoadFormFile(c, "data"); err != nil {
		return err
	}

	// Decode image
	img, err := png.Decode(bytes.NewReader(skin_b))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on create request", "data": err})
	}

	// Validate image size
	bounds := img.Bounds()

	if bounds.Max.X != 64 || bounds.Max.Y != 32 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on create request", "data": "Image have invalid size (64x32 expected)"})
	}

	// Extract head
	var head_buffer bytes.Buffer

	head_img := utils.SkinExtractHead(img)
	err = png.Encode(&head_buffer, head_img)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error on create request", "data": "Cannot extract head from image"})
	}
	head_b := head_buffer.Bytes()

	// Create entry in database
	var l = models.Skin{
		Description: input.Description,
		Public:      input.Public,
		Owner:       user,
		Data:        skin_b,
		DataHead:    head_b,
		CreatedAt:   time.Now(),
	}

	if err := database.DB.Create(&l).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Cannot interact with database", "data": err})
	}

	return c.Status(fiber.StatusOK).JSON(l)
}
