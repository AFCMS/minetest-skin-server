package middleware

import (
	"log"
	"minetest-skin-server/utils"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func AuthHandler() fiber.Handler {
	return jwtware.New(jwtware.Config{
		//Claims: jwt.RegisteredClaims{},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Println(c.Locals("user"))
			return err
		},
		SigningKey:    utils.ConfigJWTSecret,
		SigningMethod: jwt.SigningMethodHS256.Name,
		TokenLookup:   "cookie:jwt",
	})
}
