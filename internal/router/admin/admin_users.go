package admin

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"onefluid/internal/database"
	"onefluid/pgk/memcache"
	"strconv"
	"time"
)

func UpdateUser(c *fiber.Ctx) error {
	user := memcache.UserCache.Get(c.Params("username"))

	timeLocation, _ := time.LoadLocation("Europe/Moscow")
	time, _ := time.ParseInLocation("02/01/2006 15:04", c.FormValue("Subscription"), timeLocation)

	database.DB.Model(&user).Updates(map[string]interface{}{
		"Email":        c.FormValue("Email"),
		"Referrer":     c.FormValue("Referrer"),
		"HardwareID":   c.FormValue("HardwareID"),
		"Banned":       c.FormValue("Banned") == "on",
		"BannedWhy":    c.FormValue("BannedWhy"),
		"Subscription": time,
	})

	memcache.UserCaching()

	return c.Redirect("/admin/userEdit/" + user.Username)
}

func ExtendSubscriptions(c *fiber.Ctx) error {
	appendDate, _ := strconv.Atoi(c.FormValue("date"))

	database.DB.Transaction(func(tx *gorm.DB) error {
		for _, u := range memcache.UserCache {
			if err := tx.Model(&u).Update("`subscription`", u.Subscription.Add(time.Duration(appendDate)*time.Hour)).Error; err != nil {
				return err
			}
		}

		return nil
	})

	memcache.UserCaching()

	return c.Redirect("/admin/actions")
}
