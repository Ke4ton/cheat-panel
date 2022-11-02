package main

import (
	"log"
	"onefluid/internal/database"
	"onefluid/internal/router"
	"onefluid/pgk/memcache"
	"onefluid/pgk/settings"
)

func init() {
	settings.Parse()

	database.DB = database.Database()
	database.Migration()

	memcache.UserCaching()
	memcache.LogsCaching()
}

func main() {
	err := router.Application().Listen(":4000")
	if err != nil {
		log.Fatal(err.Error())
	}
}
