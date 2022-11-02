package admin

import (
	"github.com/gofiber/fiber/v2"
	"onefluid/pgk/settings"
)

func UpdateSiteSettings(c *fiber.Ctx) error {
	settings.Config.Technical = c.FormValue("technical") == "on"

	settings.Export()

	return c.Redirect("/admin/settings")
}
