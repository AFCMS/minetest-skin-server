package main

import (
	"github.com/gofiber/template/html/v2"
	"log"
	"minetest-skin-server/auth"
	"minetest-skin-server/database"
	"minetest-skin-server/routes"
	"minetest-skin-server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// Initialize Auth
	log.Println("Initializing Auth...")
	auth.Initialize()

	// Create template engine
	engine := html.New("./", ".gohtml")

	// Init Web Server
	app := fiber.New(fiber.Config{
		AppName: "Minetest Skin Server",
		Views:   engine,
	})

	// Enable CORS
	app.Use(cors.New())

	// Log requests
	app.Use(flogger.New())

	routes.SetupRoutes(app)

	log.Fatalln(app.Listen(":8080"))
}
