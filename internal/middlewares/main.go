package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"onefluid/pgk/cookiesession"
	"onefluid/pgk/language"
	"onefluid/pgk/settings"
	"strings"
	"time"
)

func IsAuthorized(c *fiber.Ctx) error {
	Authorized := c.Cookies("Authorized", "false")
	if Authorized != "true" {
		c.ClearCookie()
		return c.Redirect("/auth/sign-in")
	}

	return c.Next()
}

func AlreadyAuthorized(c *fiber.Ctx) error {
	Authorized := c.Cookies("Authorized", "false")
	if Authorized == "true" {
		return c.Redirect("/profile")
	}

	return c.Next()
}

func IsAdmin(c *fiber.Ctx) error {
	user := cookiesession.GetUser(c)
	if user.Access.Admin != true {
		return c.Redirect("/profile")
	}

	return c.Next()
}

func SiteDisabled(c *fiber.Ctx) error {
	if settings.Config.Technical == false {
		return c.Next()
	}

	if strings.Contains(c.OriginalURL(), "admin") {
		return c.Next()
	}

	return c.Render("technical", fiber.Map{})
}

func UserBanned(c *fiber.Ctx) error {
	user := cookiesession.GetUser(c)
	if user.Banned != true {
		return c.Next()
	}

	return c.Render("banned", fiber.Map{
		"BannedCause": user.BannedWhy,
	})
}

func BindData(c *fiber.Ctx) error {
	c.Bind(fiber.Map{
		"Authorized":        cookiesession.IsAuthorized(c),
		"ApplicationConfig": settings.Config,
		"CurrentTime":       time.Now(),
		"CurrentLanguage":   c.Cookies("lang", "ru"),
		"Language": func(id string) string {
			return language.Phrases[language.GetPrefix(c)+id]
		},
	})
	return c.Next()
}
