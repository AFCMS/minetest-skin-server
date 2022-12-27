package main

import (
	"log"

	"minetest-skin-server/database"
	"minetest-skin-server/middleware"
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

	// TODO: serve React frontend

	// API Routes
	api := app.Group("/api")

	api.Get("/info", routes.Info)

	// API Authentication
	api_account := api.Group("/account")

	api_account.Post("/register", routes.AccountRegister)
	api_account.Post("/login", routes.AccountLogin)
	api_account.Get("/user", middleware.AuthHandler(), routes.AccountUser)
	api_account.Post("/logout", routes.AccountLogout)

	// Interacting with skins
	api_skin := api.Group("/skin")

	api_skin.Get("/list", routes.SkinList)
	api_skin.Get("/full/:uuid", routes.SkinFull)
	api_skin.Post("/create", routes.SkinCreate)
	api_skin.Get("/recent", routes.SkinRecent)
	api_skin.Get("/rss", routes.SkinRSS)

	// Handle 404 errors
	api.All("/*", routes.NotFound)

	log.Fatalln(app.Listen(":8080"))
}
