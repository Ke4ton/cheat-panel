package api

import (
	"github.com/gofiber/fiber/v2"
	"onefluid/internal/database"
	"onefluid/internal/models"
	"onefluid/pgk/generator"
	"onefluid/pgk/memcache"
)

func CreateConfig(c *fiber.Ctx) error {
	user := memcache.UserCache.GetByHwid(c.FormValue("hwid"))
	if user.Username == "" || len(c.FormValue("name")) == 0 || len(c.FormValue("content")) == 0 {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	database.DB.Create(&models.ConfigModel{
		SecureID: generator.RandStringRunes(24),
		Name:     c.FormValue("name"),
		Author:   user.Username,
		Owner:    user.Username,
		Content:  c.FormValue("content"),
	})

	return c.JSON(fiber.Map{"Status": "Success"})
}

func CreateScript(c *fiber.Ctx) error {
	user := memcache.UserCache.GetByHwid(c.FormValue("hwid"))
	if user.Username == "" || len(c.FormValue("name")) == 0 || len(c.FormValue("content")) == 0 {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	database.DB.Create(&models.ScriptModel{
		SecureID:   generator.RandStringRunes(24),
		Name:       c.FormValue("name"),
		Author:     user.Username,
		Owner:      user.Username,
		Content:    c.FormValue("content"),
		OpenSource: true,
	})

	return c.JSON(fiber.Map{"Status": "Success"})
}

func GetConfigs(c *fiber.Ctx) error {
	user := memcache.UserCache.GetByHwid(c.FormValue("hwid"))
	if user.Username == "" {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	var userConfigs []models.ConfigModel
	database.DB.Find(&userConfigs, "`owner` = ?", user.Username)

	if len(userConfigs) == 0 {
		return c.JSON(fiber.Map{"Status": "Empty"})
	}

	return c.JSON(fiber.Map{"Status": "Success", "Configs": userConfigs})
}

func GetScripts(c *fiber.Ctx) error {
	user := memcache.UserCache.GetByHwid(c.FormValue("hwid"))
	if user.Username == "" {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	var userScripts []models.ScriptModel
	database.DB.Find(&userScripts, "`owner` = ?", user.Username)

	if len(userScripts) == 0 {
		return c.JSON(fiber.Map{"Status": "Empty"})
	}

	return c.JSON(fiber.Map{"Status": "Success", "Scripts": userScripts})
}

func SaveScript(c *fiber.Ctx) error {
	user := memcache.UserCache.GetByHwid(c.FormValue("hwid"))
	if user.Username == "" || len(c.FormValue("secureid")) == 0 || len(c.FormValue("content")) == 0 {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	var userScript models.ScriptModel
	database.DB.Find(&userScript, "`secure_id` = ?", c.FormValue("secureid"))

	if userScript.Owner != user.Username {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	database.DB.Model(&userScript).Updates(map[string]interface{}{
		"Content": c.FormValue("content"),
	})

	return c.JSON(fiber.Map{"Status": "Success"})
}

func SaveConfig(c *fiber.Ctx) error {
	user := memcache.UserCache.GetByHwid(c.FormValue("hwid"))
	if user.Username == "" || len(c.FormValue("secureid")) == 0 || len(c.FormValue("content")) == 0 {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	var userConfig models.ConfigModel
	database.DB.Find(&userConfig, "`secure_id` = ?", c.FormValue("secureid"))

	if userConfig.Owner != user.Username {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	database.DB.Model(&userConfig).Updates(map[string]interface{}{
		"Content": c.FormValue("content"),
	})

	return c.JSON(fiber.Map{"Status": "Success"})
}

func DeleteScript(c *fiber.Ctx) error {
	user := memcache.UserCache.GetByHwid(c.FormValue("hwid"))
	if user.Username == "" || len(c.FormValue("secureid")) == 0 || len(c.FormValue("content")) == 0 {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	var userScript models.ScriptModel
	database.DB.Find(&userScript, "`secure_id` = ?", c.FormValue("secureid"))

	if userScript.Owner != user.Username {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	database.DB.Delete(&userScript)

	return c.JSON(fiber.Map{"Status": "Success"})
}

func DeleteConfig(c *fiber.Ctx) error {
	user := memcache.UserCache.GetByHwid(c.FormValue("hwid"))
	if user.Username == "" || len(c.FormValue("secureid")) == 0 || len(c.FormValue("content")) == 0 {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	var userConfig models.ConfigModel
	database.DB.Find(&userConfig, "`secure_id` = ?", c.FormValue("secureid"))

	if userConfig.Owner != user.Username {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	database.DB.Delete(&userConfig)

	return c.JSON(fiber.Map{"Status": "Success"})
}
