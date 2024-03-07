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
	apiAccount := api.Group("/account")

	apiAccount.Post("/register", AccountRegister)
	apiAccount.Post("/login", AccountLogin)
	apiAccount.Get("/user", middleware.AuthHandler(), AccountUser)
	apiAccount.Post("/logout", AccountLogout)

	// Interacting with skins
	apiSkin := api.Group("/skin")

	apiSkin.Get("/list", SkinList)
	apiSkin.Get("/skin/:uuid<guid>", SkinDetails)
	apiSkin.Get("/skin/:uuid<guid>/full", SkinFull)
	apiSkin.Get("/skin/:uuid<guid>/head", SkinHead)
	apiSkin.Post("/skin/:uuid<guid>/approve", middleware.AuthHandler(), middleware.PermissionHandler(models.PermissionLevelApprover), SkinApprove)
	apiSkin.Post("/skin/:uuid<guid>/delete", middleware.AuthHandler(), NotImplemented)
	apiSkin.Post("/create", middleware.AuthHandler(), SkinCreate)
	apiSkin.Get("/recent", SkinRecent)
	apiSkin.Get("/rss", SkinRSS)

	// Interacting with users
	apiUsers := api.Group("/users")

	apiUsers.Get("/list", UsersList)
	apiUsers.Get("/list/banned", middleware.AuthHandler(), middleware.PermissionHandler(models.PermissionLevelAdmin), NotImplemented)
	apiUsers.Get("/:id<int;min(1)>", UsersID)
	apiUsers.Post("/:id<int;min(1)>/ban", middleware.AuthHandler(), middleware.PermissionHandler(models.PermissionLevelAdmin), NotImplemented)
	apiUsers.Post("/:id<int;min(1)>/unban", middleware.AuthHandler(), middleware.PermissionHandler(models.PermissionLevelAdmin), NotImplemented)
	apiUsers.Post("/:id<int;min(1)>/delete", middleware.AuthHandler(), NotImplemented)
	apiUsers.Post("/:id<int;min(1)>/permissions", middleware.AuthHandler(), middleware.PermissionHandler(models.PermissionLevelAdmin), UsersPermissions)

	// Handle 404 errors
	api.All("/*", NotFound)

	// Serve the React frontend
	app.Static("/", "./frontend/dist")
	app.Static("*", "./frontend/dist/index.html")
}
