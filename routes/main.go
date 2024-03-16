package routes

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"

	"minetest-skin-server/middleware"
	"minetest-skin-server/models"
	"minetest-skin-server/utils"
)

func SetupRoutes(app *fiber.App) {
	// API Routes
	api := app.Group("/api")

	api.Get("/info", Info)

	// API Authentication
	apiAccount := api.Group("/account")

	apiAccount.Post("/register", AccountRegister)
	apiAccount.Post("/login", AccountLogin)
	apiAccount.Get("/user", middleware.AuthHandler, AccountUser)
	apiAccount.Post("/logout", AccountLogout)

	// Interacting with skins
	apiSkin := api.Group("/skin")

	apiSkin.Get("/list", SkinList)
	apiSkin.Get("/skin/:uuid<guid>", SkinDetails)
	apiSkin.Get("/skin/:uuid<guid>/full", SkinFull)
	apiSkin.Get("/skin/:uuid<guid>/head", SkinHead)
	apiSkin.Post("/skin/:uuid<guid>/approve", middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelApprover), SkinApprove)
	apiSkin.Post("/skin/:uuid<guid>/delete", middleware.AuthHandler, NotImplemented)
	apiSkin.Post("/create", middleware.AuthHandler, SkinCreate)
	apiSkin.Get("/recent", SkinRecent)
	apiSkin.Get("/rss", SkinRSS)

	// Interacting with users
	apiUsers := api.Group("/users")

	apiUsers.Get("/list", UsersList)
	apiUsers.Get("/list/banned", middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelAdmin), NotImplemented)
	apiUsers.Get("/:id<int;min(1)>", UsersID)
	apiUsers.Post("/:id<int;min(1)>/ban", middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelAdmin), NotImplemented)
	apiUsers.Post("/:id<int;min(1)>/unban", middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelAdmin), NotImplemented)
	apiUsers.Post("/:id<int;min(1)>/delete", middleware.AuthHandler, NotImplemented)
	apiUsers.Post("/:id<int;min(1)>/permissions", middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelAdmin), UsersPermissions)

	// Handle 404 errors
	api.All("*", NotFound)

	// Serve the React frontend
	if utils.ConfigFrontendDevMode {
		app.Get("*", proxy.Balancer(proxy.Config{
			Servers: []string{utils.ConfigFrontendURL},
			ModifyResponse: func(c *fiber.Ctx) error {
				if c.Response().StatusCode() == fiber.StatusNotFound {
					return c.Status(fiber.StatusOK).Render("index", fiber.Map{
						"DevMode": utils.ConfigFrontendDevMode,
					})
				}
				return nil
			},
		}))
	} else {
		// Parse JSON Vite manifest
		manifest := utils.ViteManifest{}
		data, err := os.ReadFile("./frontend/dist/.vite/manifest.json")
		if err != nil {
			log.Fatal(err)
		}
		if err == nil {
			err = json.Unmarshal(data, &manifest)
		}

		app.Static("/", "./frontend/dist")
		app.Get("*", func(c *fiber.Ctx) error {
			return c.Render("index", fiber.Map{
				"DevMode": false,
				"MainCSS": manifest["src/main.tsx"].Css[0],
				"MainJS":  manifest["src/main.tsx"].File,
			})
		})
	}
}
