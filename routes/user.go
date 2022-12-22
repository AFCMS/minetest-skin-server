package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthenticated"})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var user models.Account

	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"version": "1.0"})
}
