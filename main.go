package main

import (
	"log"

	"minetest-skin-server/database"
	"minetest-skin-server/routes"
	"minetest-skin-server/utils"

	"github.com/gofiber/fiber/v2"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Check for OptiPNG installation
	if utils.ConfigOptipngEnabled {
		optipngPresent := utils.OptiPNGPresent()
		if optipngPresent {
			log.Println("OptiPNG found")
		} else {
			log.Fatalln("OptiPNG not found")
		}
	}

	// Connection to Database
	log.Println("Connecting to Database...")

	database.ConnectDB()

	// Init Web Server
	app := fiber.New(fiber.Config{
		AppName: "Minetest Skin Server",
	})

	app.Use(flogger.New())

	routes.SetupRoutes(app)

	log.Fatalln(app.Listen(":8080"))
}
