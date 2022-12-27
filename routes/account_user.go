package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AccountUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)

	// TODO: use RegisteredClaims instead of MapClaims
	claims := token.Claims.(jwt.MapClaims)

	var user models.Account

	cs, err := strconv.ParseInt(claims["iss"].(string), 10, 0)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	user, err = database.AccountFromID(int(cs))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":               user.ID,
		"name":             user.Name,
		"email":            user.Email,
		"permission_level": user.PermissionLevel,
	})
}
