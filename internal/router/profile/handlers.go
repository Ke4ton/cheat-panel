package profile

import (
	"github.com/gofiber/fiber/v2"
	"onefluid/pgk/cookiesession"
)

func IndexPage(c *fiber.Ctx) error {
	referrer := c.Query("ref", "none")
	if referrer != "none" {
		c.Cookie(&fiber.Cookie{Name: "ref", Value: referrer})
	}

	return c.Render("index", fiber.Map{
		"User": cookiesession.GetUser(c),
	})
}

func SignUpPage(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{
		"Error": c.Query("err"),
	})
}

func SignInPage(c *fiber.Ctx) error {
	return c.Render("signin", fiber.Map{
		"Error": c.Query("err"),
	})
}

func ProfilePage(c *fiber.Ctx) error {
	user := cookiesession.GetUser(c)

	return c.Render("profile", fiber.Map{
		"User":         user,
		"Error":        c.Query("err"),
		"Subscription": user.Subscription.Format("2006/01/02 15:04"),
	})
}
