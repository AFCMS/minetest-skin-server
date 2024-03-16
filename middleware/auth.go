package middleware

import (
	"github.com/gofiber/fiber/v2"
	"minetest-skin-server/auth"
	"minetest-skin-server/database"
)

// AuthHandler Check if the user is authenticated
//
// Put the database entry for the user in locals
func AuthHandler(c *fiber.Ctx) error {
	sess, err := auth.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	userAccount, err := database.AccountFromID(sess.Get("uid").(uint))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	c.Locals("session", sess)
	c.Locals("user", userAccount)
	return c.Next()
}
