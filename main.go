package main

import (
	"log"

	"minetest-skin-server/database"
	"minetest-skin-server/routes"

	"github.com/gofiber/fiber/v2"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Reading Env file...")
	_ = godotenv.Load()

	// Connection to Database
	log.Println("Connecting to Database...")

	database.ConnectDB()

	// Init Web Server
	app := fiber.New(fiber.Config{
		AppName: "Minetest Skin Server",
	})

	app.Use(flogger.New())

	// API Routes

	api := app.Group("/api")

	api.Get("/info", routes.Info)
	api.Post("/register", routes.Register)
	api.Post("/login", routes.Login)
	api.Get("/user", routes.User)
	api.Post("/logout", routes.Logout)

	log.Fatalln(app.Listen(":8080"))
}
