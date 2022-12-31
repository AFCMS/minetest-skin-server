package middleware

import (
	"minetest-skin-server/database"
	"minetest-skin-server/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

// Check if the user is authenticated
//
// Put the database entry for the user in locals
func AuthHandler() fiber.Handler {
	return jwtware.New(jwtware.Config{
		//Claims: jwt.RegisteredClaims{},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid JWT", "data": err})
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals("user_jwt").(*jwt.Token)

			// TODO: use RegisteredClaims instead of MapClaims
			claims := token.Claims.(jwt.MapClaims)

			cs, err := strconv.ParseInt(claims["iss"].(string), 10, 0)

			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid JWT"})
			}

			user_account, err := database.AccountFromID(int(cs))

			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
			}

			c.Locals("user", user_account)
			return c.Next()
		},
		ContextKey:    "user_jwt",
		SigningKey:    utils.ConfigJWTSecret,
		SigningMethod: jwt.SigningMethodHS256.Name,
		TokenLookup:   "cookie:jwt",
	})
}
