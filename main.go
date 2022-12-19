package main

import (
	"log"

	"minetest-skin-server/database"
	"minetest-skin-server/routes"

	"github.com/gofiber/fiber/v2"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Connection to Database
	log.Println("Connecting to Database...")

	database.ConnectDB()

	// Init Web Server
	app := fiber.New(fiber.Config{
		AppName: "Minetest Skin Server",
	})

	app.Use(flogger.New())

	api := app.Group("/api")

	api.Post("/login", routes.Login)

	log.Fatalln(app.Listen(":8080"))
}
