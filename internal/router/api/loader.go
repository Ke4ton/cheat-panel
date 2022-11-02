package api

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
	"onefluid/internal/database"
	"onefluid/internal/models"
	"onefluid/pgk/logger"
	"time"
)

func Auth(c *fiber.Ctx) error {
	username, password := c.FormValue("Username"), c.FormValue("Password")

	if c.FormValue("HardwareID") == "" ||
		c.FormValue("Username") == "" ||
		c.FormValue("Password") == "" {
		return c.JSON(fiber.Map{"Status": "InvalidPacket"})
	}

	var findUser models.UserModel
	database.DB.Find(&findUser, "`username` = ?", username)
	if findUser.Username != username {
		return c.JSON(fiber.Map{"Status": "UserNotFound"})
	}

	passwordHash := md5.Sum([]byte(password))
	hashString := hex.EncodeToString(passwordHash[:])
	if hashString != findUser.Password {
		return c.JSON(fiber.Map{"Status": "WrongPassword"})
	}

	if findUser.HardwareID == "" {
		database.DB.Model(&findUser).Updates(map[string]interface{}{
			"HardwareID": c.FormValue("HardwareID"),
		})
	}

	if findUser.HardwareID != c.FormValue("HardwareID") {
		return c.JSON(fiber.Map{"Status": "HardwareError"})
	}

	if findUser.Banned == true {
		return c.JSON(fiber.Map{"Status": "Suspended"})
	}

	if findUser.Subscription.Unix() < time.Now().Unix() {
		return c.JSON(fiber.Map{"Status": "LicenseNotFound"})
	}

	logger.Database(findUser.Username, "Injected cheat", logger.Injection, c)

	return c.JSON(fiber.Map{
		"Status": "Success",
		"User":   findUser,
	})
}

func Telegram(c *fiber.Ctx) error {
	logger.Telegram(c.FormValue("Message"))
	return c.JSON(fiber.Map{"Status": "OK"})
}
