package middleware

import (
	"github.com/gofiber/fiber/v3"

	"minetest-skin-server/auth"
	"minetest-skin-server/database"
)

// AuthHandler Check if the user is authenticated
//
// Put the database entry for the user in locals
func AuthHandler(c fiber.Ctx) error {
	sess, err := auth.SessionStore.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if v, ok := sess.Get("uid").(uint); !ok || v == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Not logged in"})
	}

	userAccount, err := database.AccountFromID(sess.Get("uid").(uint))

	if err != nil {
		err := sess.Destroy()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	c.Locals("session", sess)
	c.Locals("user", userAccount)
	return c.Next()
}
