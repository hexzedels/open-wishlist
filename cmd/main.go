package main

import (
	"openwishlist/app/core"
	"openwishlist/app/db"
	"openwishlist/app/run"
)

func main() {
	run.Init()

	postgresClient := db.NewPostgresClient()

	bot := core.NewTelegramBot(postgresClient)
	bot.Start()
}
