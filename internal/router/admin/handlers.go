package admin

import (
	"github.com/gofiber/fiber/v2"
	"onefluid/internal/database"
	"onefluid/internal/models"
	"onefluid/pgk/cookiesession"
	"onefluid/pgk/memcache"
)

func AdminStatistic(c *fiber.Ctx) error {
	return c.Render("admin/admin", fiber.Map{
		"User":              cookiesession.GetUser(c),
		"Registered":        len(memcache.UserCache),
		"LumenSubscribed":   memcache.UserCache.WithSubscriptions(),
		"LumenUnsubscribed": memcache.UserCache.WithoutSubscriptions(),

		"CountInjections":     memcache.LogsCache.CountInjections(),
		"CountAuthorizations": memcache.LogsCache.CountAuthorizations(),
		"CountRegistrations":  memcache.LogsCache.CountRegistrations(),
	})
}

func AdminFastActions(c *fiber.Ctx) error {
	return c.Render("admin/admin-fast-actions", fiber.Map{
		"User": cookiesession.GetUser(c),
	})
}

func AdminSettings(c *fiber.Ctx) error {
	return c.Render("admin/admin-settings", fiber.Map{
		"User": cookiesession.GetUser(c),
	})
}

func AdminUsers(c *fiber.Ctx) error {
	return c.Render("admin/admin-user", fiber.Map{
		"User":  cookiesession.GetUser(c),
		"Users": memcache.UserCache,
	})
}

func AdminBannedUsers(c *fiber.Ctx) error {
	return c.Render("admin/admin-banned-user", fiber.Map{
		"User":  cookiesession.GetUser(c),
		"Users": memcache.UserCache.Banned(),
	})
}

func AdminEditUser(c *fiber.Ctx) error {
	editUser := memcache.UserCache.Get(c.Params("username"))
	if editUser.Username == "" {
		return c.Redirect("/admin/user")
	}

	var UserLogs []models.UserLog
	database.DB.Order("ID DESC").Limit(50).Find(&UserLogs, "`username` = ?", editUser.Username)

	return c.Render("admin/admin-edit-user", fiber.Map{
		"User":     cookiesession.GetUser(c),
		"UserLogs": UserLogs,
		"EditUser": memcache.UserCache.Get(c.Params("username")),
	})
}

func AdminPromos(c *fiber.Ctx) error {
	var buckets []models.PromoBucket
	database.DB.Find(&buckets)

	return c.Render("admin/admin-promo", fiber.Map{
		"User":    cookiesession.GetUser(c),
		"Folders": buckets,
	})
}

func AdminPromoBucket(c *fiber.Ctx) error {
	var buckets models.PromoBucket
	database.DB.Find(&buckets, "id = ?", c.Params("id"))

	var keys []models.PromoCode
	database.DB.Find(&keys, "`bucket_id` = ?", buckets.ID)

	return c.Render("admin/admin-promo-bucket", fiber.Map{
		"User":   cookiesession.GetUser(c),
		"Folder": buckets,
		"Keys":   keys,
	})
}
