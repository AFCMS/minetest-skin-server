package routes

import (
	"minetest-skin-server/database"
	"minetest-skin-server/models"
	"minetest-skin-server/types"
	"minetest-skin-server/utils"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	input := types.InputRegister{}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on register request", "data": err.Error()})
	}

	if utils.IsValidEmail(input.Email) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on register request", "data": "Invalid email"})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	user := models.Account{
		Name:         input.Name,
		Email:        input.Email,
		Password:     password,
		CreationDate: time.Now().Unix(),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on register request", "data": err.Error()})
	}

	return c.JSON(types.OutputRegister{
		Id:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		CreationDate: user.CreationDate,
	})
}
