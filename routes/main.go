package routes

import (
	"minetest-skin-server/middleware"
	"minetest-skin-server/models"

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
	api_skin.Get("/skin/:uuid", SkinDetails)
	api_skin.Get("/skin/:uuid/full", SkinFull)
	api_skin.Get("/skin/:uuid/head", SkinHead)
	api_skin.Post("/skin/:uuid/approve", middleware.AuthHandler(), middleware.PermissionHandler(models.PermissionLevelApprover), SkinApprove)
	api_skin.Post("/skin/:uuid/delete", middleware.AuthHandler(), NotImplemented)
	api_skin.Post("/create", middleware.AuthHandler(), SkinCreate)
	api_skin.Get("/recent", SkinRecent)
	api_skin.Get("/rss", SkinRSS)

	// Interacting with users
	api_users := api.Group("/users")

	api_users.Get("/list", UsersList)
	api_users.Get("/list/banned", middleware.AuthHandler(), middleware.PermissionHandler(models.PermissionLevelAdmin), NotImplemented)
	api_users.Get("/:id", UsersID)
	api_users.Post("/:id/ban", middleware.AuthHandler(), middleware.PermissionHandler(models.PermissionLevelAdmin), NotImplemented)
	api_users.Post("/:id/unban", middleware.AuthHandler(), middleware.PermissionHandler(models.PermissionLevelAdmin), NotImplemented)
	api_users.Post("/:id/delete", middleware.AuthHandler(), NotImplemented)
	api_users.Post("/:id/permissions", middleware.AuthHandler(), middleware.PermissionHandler(models.PermissionLevelAdmin), UsersPermissions)

	// Handle 404 errors
	api.All("/*", NotFound)

	// Serve the React frontend
	app.Static("/", "./frontend/dist")
	app.Static("*", "./frontend/dist/index.html")
}
