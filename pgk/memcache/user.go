package memcache

import (
	"onefluid/internal/database"
	"onefluid/internal/models"
	"time"
)

type userArray []models.UserModel

var UserCache userArray

func UserCaching() {
	database.DB.Order("'created_at' DESC").Find(&UserCache)
}

func (c userArray) Get(findUser string) models.UserModel {
	for _, u := range c {
		if u.Username == findUser {
			return u
		}
	}

	return models.UserModel{}
}

func (c userArray) GetByHwid(hwidUser string) models.UserModel {
	for _, u := range c {
		if u.HardwareID == hwidUser {
			return u
		}
	}

	return models.UserModel{}
}

func (c userArray) GetReferrals(targetUser string) (u []models.UserModel) {
	for _, d := range c {
		if d.Referrer == targetUser {
			u = append(u, d)
		}
	}

	return
}

func (c userArray) WithSubscriptions() (total int) {
	for _, d := range c {
		if d.Subscription.Unix() > time.Now().Unix() {
			total += 1
		}
	}
	return
}

func (c userArray) WithoutSubscriptions() (total int) {
	for _, d := range c {
		if d.Subscription.Unix() < time.Now().Unix() {
			total += 1
		}
	}
	return
}

func (c userArray) Banned() (a []models.UserModel) {
	for _, u := range c {
		if u.Banned {
			a = append(a, u)
		}
	}

	return
}
