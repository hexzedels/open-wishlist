package main

import (
	"context"
	"os"

	"openwishlist/app/core"
	"openwishlist/app/db"
	"openwishlist/app/run"
	"openwishlist/app/sdk"
)

func main() {
	run.Init()

	postgresClient := db.NewPostgresClient(context.Background(), os.Getenv(sdk.EnvPostgres))

	logger := run.Logger.Named("telegram bot")

	bot := core.NewTelegramBot(postgresClient, logger)
	bot.Start()
}
