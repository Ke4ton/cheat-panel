package logger

import (
	"github.com/gofiber/fiber/v2"
	goip "github.com/jpiontek/go-ip-api"
	"net/http"
	"net/url"
	"onefluid/internal/database"
	"onefluid/internal/models"
	"onefluid/pgk/memcache"
	"onefluid/pgk/settings"
	"time"
)

const (
	Injection = iota
	Authorizations
	Registrations
	Other
)

func Database(username, message string, log int, ctx *fiber.Ctx) {
	var ip string

	if settings.Config.Debug == true {
		ip = "195.24.111.111"
	} else {
		ip = ctx.IPs()[0]
	}

	client := goip.NewClient()
	location, _ := client.GetLocationForIp(ip)

	database.DB.Create(&models.UserLog{
		Username: username,
		Message:  message,
		Country:  location.Country,
		City:     location.City,
		IP:       ip,
		Type:     log,
		Day:      int(time.Now().Unix() / 86400),
	})

	go memcache.LogsCaching()
}

func Telegram(message string) {
	data := url.Values{
		"chat_id": {settings.Config.TelegramID},
		"text":    {message},
	}

	http.PostForm("https://api.telegram.org/bot"+settings.Config.Telegram+"/sendMessage", data)
}
