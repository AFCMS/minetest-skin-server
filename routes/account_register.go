package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"
	"minetest-skin-server/types"
	"minetest-skin-server/utils"
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

func AccountRegister(c *fiber.Ctx) error {
	input := types.InputRegister{}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on register request", "data": err.Error()})
	}

	if !utils.IsValidEmail(input.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on register request", "data": "Invalid email"})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	// TODO: validate email
	var parsedEmail *mail.Address
	var err error
	if parsedEmail, err = mail.ParseAddress(input.Email); err != nil || parsedEmail.Name != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on email format", "data": "Invalid email"})
	}

	user := models.Account{
		Name:      input.Name,
		Email:     parsedEmail.Address,
		Password:  password,
		CreatedAt: time.Now(),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on register request", "data": err.Error()})
	}

	return c.JSON(types.OutputRegister{
		Id:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		CreationDate: user.CreatedAt.Unix(),
	})
}
