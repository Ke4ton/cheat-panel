package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"onefluid/internal/models"
	"onefluid/pgk/settings"
)

var DB *gorm.DB

func Database() *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(settings.Config.DSN),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func Migration() {
	err := DB.AutoMigrate(
		&models.UserModel{},
		&models.UserLog{},

		&models.PromoBucket{},
		&models.PromoCode{},
		&models.PromoActivations{},

		&models.ConfigModel{},
		&models.ScriptModel{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}
}
