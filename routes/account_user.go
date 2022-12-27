package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AccountUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)

	// TODO: use RegisteredClaims instead of MapClaims
	claims := token.Claims.(jwt.MapClaims)

	var user models.Account

	if err := database.DB.Where("id = ?", claims["iss"]).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":               user.ID,
		"name":             user.Name,
		"email":            user.Email,
		"permission_level": user.PermissionLevel,
	})
}
