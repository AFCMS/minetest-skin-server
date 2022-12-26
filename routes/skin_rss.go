package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SkinRSS(c *fiber.Ctx) error {
	c.Type("application/rss+xml")

	return c.XML(fiber.Map{
		"note": fiber.Map{
			"language": "en",
			"title":    "Some event",
		},
	})
}
