package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"onefluid/internal/middlewares"
	"onefluid/internal/router/admin"
	"onefluid/internal/router/api"
	"onefluid/internal/router/profile"
	"onefluid/pgk/settings"
	"time"
)

func Application() (app *fiber.App) {
	viewEngine := html.New("./web/templates", ".html")
	viewEngine.Reload(settings.Config.Debug)

	app = fiber.New(
		fiber.Config{Views: viewEngine},
	)

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: "NAS3RFUxkhjAKg3ANv6tznvbTUBxw0sC",
	}))

	app.Use(logger.New(logger.Config{
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
	}))

	app.Static("/assets", "./web/static")

	app.Use(middlewares.BindData)
	app.Use(middlewares.SiteDisabled)
	app.Use(middlewares.UserBanned)

	app.Get("/", profile.IndexPage)

	app.Get("/lang", profile.SetLang)

	app.Get("/logout", profile.LogOut)
	auth := app.Group("/auth", middlewares.AlreadyAuthorized)
	auth.Get("/sign-up", profile.SignUpPage)
	auth.Get("/sign-in", profile.SignInPage)
	auth.Post("/actions/sign-up", profile.SignUp)
	auth.Post("/actions/sign-in", profile.SignIn)

	app.Get("/activate/:key", middlewares.IsAuthorized, profile.ActivateKey)

	profilePage := app.Group("/profile", middlewares.IsAuthorized)
	profilePage.Get("/", profile.ProfilePage)
	profilePage.Post("/actions/activatePromo", profile.ActivateKeyForm)

	loaderApi := app.Group("/api/loader")
	loaderApi.Post("/authLoader", api.Auth)
	loaderApi.Post("/telegramLog", api.Telegram)

	cloudApi := app.Group("/api/cloud")
	cloudApi.Post("/createConfig", api.CreateConfig)
	cloudApi.Post("/createScript", api.CreateScript)
	cloudApi.Post("/getConfigs", api.GetConfigs)
	cloudApi.Post("/getScripts", api.GetScripts)
	cloudApi.Post("/saveConfig", api.SaveConfig)
	cloudApi.Post("/saveScript", api.SaveScript)
	cloudApi.Post("/deleteConfig", api.DeleteConfig)
	cloudApi.Post("/deleteScript", api.DeleteScript)

	adminGroup := app.Group("/admin", middlewares.IsAuthorized, middlewares.IsAdmin)
	adminGroup.Get("/", admin.AdminStatistic)
	adminGroup.Get("/user", admin.AdminUsers)
	adminGroup.Get("/promo", admin.AdminPromos)
	adminGroup.Get("/actions", admin.AdminFastActions)
	adminGroup.Get("/settings", admin.AdminSettings)
	adminGroup.Get("/bannedUsers", admin.AdminBannedUsers)
	adminGroup.Get("/promo/bucket/:id", admin.AdminPromoBucket)
	adminGroup.Get("/userEdit/:username", admin.AdminEditUser)
	adminGroup.Get("/actions/exportBucket/:id", admin.ExportBucket)
	adminGroup.Get("/actions/deleteKey/:code/:bucket", admin.DeleteKey)
	adminGroup.Post("/actions/updateUser/:username", admin.UpdateUser)
	adminGroup.Post("/actions/extendSubscription", admin.ExtendSubscriptions)
	adminGroup.Post("/actions/updateSiteSettings", admin.UpdateSiteSettings)
	adminGroup.Post("/actions/createBucket", admin.CreateBucket)
	adminGroup.Post("/actions/generateRandom", admin.GenerateRandomPromo)
	adminGroup.Post("/actions/generatePromo", admin.GeneratePromo)

	return app
}
