package models

import (
	"gorm.io/gorm"
	"time"
)

type AdminAccess struct {
	Admin    bool
	Actions  bool
	Billings bool
	Promo    bool
	Users    bool

	PriceSettings   bool
	SiteSettings    bool
	ProjectSettings bool

	SetAdmins bool
}

type UserModel struct {
	gorm.Model

	Username     string `gorm:"unique"`
	Password     string
	Email        string `gorm:"unique"`
	Referrer     string
	HardwareID   string
	Banned       bool
	BannedWhy    string
	Subscription time.Time

	Access AdminAccess `gorm:"embedded;embeddedPrefix:admin_"`
}

type UserLog struct {
	gorm.Model

	Username string
	Message  string
	Country  string
	City     string
	IP       string

	Type int
	Day  int
}
