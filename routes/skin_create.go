package routes

import (
	"bytes"
	"image/png"
	"minetest-skin-server/database"
	"minetest-skin-server/models"
	"minetest-skin-server/types"
	"minetest-skin-server/utils"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// SkinCreate Handle Skin creation
//
// Use a multipart request
func SkinCreate(c *fiber.Ctx) error {
	// Get User
	user := c.Locals("user").(models.Account)

	input := types.InputSkinCreate{}

	// Get the text fields
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorOutput{Message: "Invalid request body", Data: err.Error()})
	}

	// Get file part
	var skin_b []byte
	var err error
	if skin_b, err = utils.LoadFormFile(c, "data"); err != nil {
		return err
	}

	// Decode image
	img, err := png.Decode(bytes.NewReader(skin_b))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorOutput{Message: "Cannot decode skin", Data: err.Error()})
	}

	// Validate image size
	bounds := img.Bounds()

	if bounds.Max.X != 64 || bounds.Max.Y != 32 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorOutput{Message: "Invalid skin", Data: "Image have invalid size (64x32 expected)"})
	}

	// Extract head
	var head_buffer bytes.Buffer

	head_img := utils.SkinExtractHead(img)
	err = png.Encode(&head_buffer, head_img)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorOutput{Message: "Server error", Data: "Cannot extract head from image"})
	}
	head_b := head_buffer.Bytes()

	var skin_b_opti = skin_b
	var head_b_opti = head_b

	// Optionally run OptiPNG
	// TODO: Run them async to get them done faster
	// https://stackoverflow.com/questions/27792389/golang-functions-parallel-execution-with-return
	if utils.ConfigOptipngEnabled {
		var out_1 []byte
		var err_1 error

		var out_2 []byte
		var err_2 error

		var sg sync.WaitGroup

		sg.Add(2)

		go func(out *[]byte, err *error, sg *sync.WaitGroup) {
			*out, *err = utils.OptiPNGBytes(skin_b)
			sg.Done()
		}(&out_1, &err_1, &sg)

		go func(out *[]byte, err *error, sg *sync.WaitGroup) {
			*out, *err = utils.OptiPNGBytes(head_b)
			sg.Done()
		}(&out_2, &err_2, &sg)

		sg.Wait()

		if err_1 != nil || err_2 != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorOutput{Message: "Server error", Data: "Cannot obtimize image"})
		}

		skin_b_opti = out_1
		head_b_opti = out_2
	}

	// Create entry in database
	var l = models.Skin{
		Description: input.Description,
		Public:      input.Public,
		Owner:       user,
		Data:        skin_b_opti,
		DataHead:    head_b_opti,
		CreatedAt:   time.Now(),
	}

	if err := database.DB.Create(&l).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorOutput{Message: "Cannot interact with database", Data: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(l)
}
