package main

import (
	"openwishlist/app/core"
	"openwishlist/app/db"
	"openwishlist/app/run"
)

func main() {
	run.Init()

	postgresClient := db.NewPostgresClient()

	logger := run.Logger.Named("telegram bot")

	bot := core.NewTelegramBot(postgresClient, logger)
	bot.Start()
}
