package routes

import (
	"minetest-skin-server/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// TODO: serve React frontend

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
	api_skin.Get("/full/:uuid", SkinFull)
	api_skin.Post("/create", SkinCreate)
	api_skin.Get("/recent", SkinRecent)
	api_skin.Get("/rss", SkinRSS)

	// Handle 404 errors
	api.All("/*", NotFound)
}
