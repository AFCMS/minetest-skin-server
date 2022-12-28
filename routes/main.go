package routes

import (
	"minetest-skin-server/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// API Routes
	api := app.Group("/api")

	api.Get("/info", Info)

	// API Authentication
	api_account := api.Group("/account")

	api_account.Post("/register", AccountRegister)
	api_account.Post("/login", AccountLogin)
	api_account.Get("/user", middleware.AuthHandler(), AccountUser)
	api_account.Post("/logout", AccountLogout)

	// Interacting with skins
	api_skin := api.Group("/skin")

	api_skin.Get("/list", SkinList)
	api_skin.Get("/full/:uuid", NotImplemented)
	api_skin.Get("/full/:uuid/full", SkinFull)
	api_skin.Get("/full/:uuid/head", NotImplemented)
	api_skin.Post("/create", SkinCreate)
	api_skin.Get("/recent", SkinRecent)
	api_skin.Get("/rss", SkinRSS)

	// Handle 404 errors
	api.All("/*", NotFound)

	// Serve the React frontend
	app.Static("/", "./frontend/build")
	app.Static("*", "./frontend/build/index.html")
}
