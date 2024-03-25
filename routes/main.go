package routes

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/proxy"

	"minetest-skin-server/auth"
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
	apiAccount.Get("/user", AccountUser, middleware.AuthHandler)
	apiAccount.Post("/logout", AccountLogout, middleware.AuthHandler)

	apiOauthEndpoints := apiAccount.Group("/providers")

	auth.RegisterEndpoints(apiOauthEndpoints)

	// Interacting with skins
	apiSkin := api.Group("/skin")

	apiSkin.Get("/list", SkinList)
	apiSkin.Get("/skin/:uuid<guid>", SkinDetails)
	apiSkin.Get("/skin/:uuid<guid>/full", SkinFull)
	apiSkin.Get("/skin/:uuid<guid>/head", SkinHead)
	apiSkin.Post("/skin/:uuid<guid>/approve", SkinApprove, middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelApprover))
	apiSkin.Post("/skin/:uuid<guid>/delete", NotImplemented, middleware.AuthHandler)
	apiSkin.Post("/create", SkinCreate, middleware.AuthHandler)
	apiSkin.Get("/recent", SkinRecent)
	apiSkin.Get("/rss", SkinRSS)

	// Interacting with users
	apiUsers := api.Group("/users")

	apiUsers.Get("/list", UsersList)
	apiUsers.Get("/list/banned", NotImplemented, middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelAdmin))
	apiUsers.Get("/:id<int;min(1)>", UsersID)
	apiUsers.Post("/:id<int;min(1)>/ban", NotImplemented, middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelAdmin))
	apiUsers.Post("/:id<int;min(1)>/unban", NotImplemented, middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelAdmin))
	apiUsers.Post("/:id<int;min(1)>/delete", NotImplemented, middleware.AuthHandler)
	apiUsers.Post("/:id<int;min(1)>/permissions", UsersPermissions, middleware.AuthHandler, middleware.PermissionHandler(models.PermissionLevelAdmin))

	// Handle 404 errors
	api.All("*", NotFound)

	// Serve the React frontend
	if utils.ConfigFrontendDevMode {
		app.Get("*", proxy.Balancer(proxy.Config{
			Servers: []string{utils.ConfigFrontendURL},
			ModifyResponse: func(c fiber.Ctx) error {
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
		err = json.Unmarshal(data, &manifest)
		if err != nil {
			log.Fatal(err)
		}

		app.Static("/", "./frontend/dist")
		app.Get("*", func(c fiber.Ctx) error {
			return c.Render("index", fiber.Map{
				"DevMode":                false,
				"MainCSS":                manifest["src/main.tsx"].Css[0],
				"MainJS":                 manifest["src/main.tsx"].File,
				"GoogleSiteVerification": utils.ConfigVerificationGoogle,
			})
		})
	}
}
