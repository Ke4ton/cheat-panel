package cookiesession

import (
	"github.com/gofiber/fiber/v2"
	"onefluid/internal/models"
	"onefluid/pgk/memcache"
)

func GetUser(c *fiber.Ctx) models.UserModel {
	username := c.Cookies("USR", "JDIAHJJFAIHADWJH213LHAD891H23L")
	if IsAuthorized(c) == true {
		if memcache.UserCache.Get(username).Username == "" {
			c.Redirect("/auth/sign-in")
			return models.UserModel{}
		}
	}

	return memcache.UserCache.Get(username)
}

func IsAuthorized(c *fiber.Ctx) bool {
	Authorized := c.Cookies("Authorized", "false")
	if Authorized == "true" {
		return true
	}
	return false
}
