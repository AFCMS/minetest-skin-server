package utils

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

func LoadFormFile(c *fiber.Ctx, name string) ([]byte, error) {
	// Get file header
	var file_header *multipart.FileHeader
	var err error
	if file_header, err = c.FormFile("data"); err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on create request", "data": err})
	}

	// Validate file size
	if file_header.Size > 100000 {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on create request", "data": "Too big file"})
	}

	// Get file
	var file multipart.File
	if file, err = file_header.Open(); err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on create request", "data": err})
	}

	var b = make([]byte, file_header.Size)

	// Read file content
	var _ int
	if _, err = file.Read(b); err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on create request", "data": err})
	}

	return b, nil
}
