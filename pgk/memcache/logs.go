package memcache

import (
	"onefluid/internal/database"
	"onefluid/internal/models"
	"time"
)

type logsArray []models.UserLog

var LogsCache logsArray

func LogsCaching() {
	database.DB.Where("day >= ? and (type = 0 or type = 1 or type = 2)", (time.Now().Unix()/86400)-7).
		Find(&LogsCache)
}

func (u logsArray) CountInjections() (c [8]int) {
	for _, d := range u {
		day := 7 - (int(time.Now().Unix()/86400) - d.Day)
		if d.Type == 0 {
			c[day] += 1
		}
	}
	return
}

func (u logsArray) CountAuthorizations() (c [8]int) {
	for _, d := range u {
		day := 7 - (int(time.Now().Unix()/86400) - d.Day)
		if d.Type == 1 {
			c[day] += 1
		}
	}

	return
}

func (u logsArray) CountRegistrations() (c [8]int) {
	for _, d := range u {
		day := 7 - (int(time.Now().Unix()/86400) - d.Day)
		if d.Type == 2 {
			c[day] += 1
		}
	}

	return
}
