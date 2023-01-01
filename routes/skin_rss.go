package routes

import (
	"github.com/gofiber/fiber/v2"
)

type entryDetails struct {
	Language string `xml:"language"`
	Title    string `xml:"title"`
}

type entry struct {
	Note entryDetails `xml:"note"`
}

func SkinRSS(c *fiber.Ctx) error {
	err := c.XML(entry{
		Note: entryDetails{
			Language: "en",
			Title:    "Test",
		},
	})

	if err != nil {
		return err
	}

	c.Type("application/rss+xml")
	return nil
}
