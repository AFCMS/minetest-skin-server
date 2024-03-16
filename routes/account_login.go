package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"minetest-skin-server/auth"
	"minetest-skin-server/database"
	"minetest-skin-server/models"
	"minetest-skin-server/types"
	"time"
)

func AccountLogin(c *fiber.Ctx) error {
	input := types.InputLogin{}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on login request", "data": err})
	}

	var user models.Account

	// Find user by email

	user, err := database.AccountFromEmail(input.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	if user.Banned {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "User is banned", "reason": user.BanReason})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Incorrect password"})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not login"})
	}

	sess, err := auth.SessionStore.Get(c)

	if sess.Fresh() {
		// Get session ID
		sid := sess.ID()

		// Get user ID
		uid := user.ID

		// Save session data
		sess.Set("uid", uid)
		sess.Set("sid", sid)
		sess.Set("ip", c.Context().RemoteIP().String())
		sess.Set("login", time.Unix(time.Now().Unix(), 0).UTC().String())
		sess.Set("ua", string(c.Request().Header.UserAgent()))

		err := sess.Save()
		if err != nil {
			log.Println(err)
		}
	}

	if err := database.AccountSetLastConnection(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error interacting with database", "data": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Success"})
}
