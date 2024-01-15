package core

import (
	"os"

	"openwishlist/app/db"
	"openwishlist/app/sdk"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type TelegramBot struct {
	dbClient db.IClient
	bot      *tgbotapi.BotAPI
	logger   *zap.Logger
}

func NewTelegramBot(dbClient db.IClient, logger *zap.Logger) *TelegramBot {
	bot, err := tgbotapi.NewBotAPI(os.Getenv(sdk.EnvToken))
	if err != nil {
		panic(err)
	}

	return &TelegramBot{
		dbClient: dbClient,
		bot:      bot,
	}
}

func (r *TelegramBot) Start() {
	defer func() {
		r.logger.Info("stopped telegram bot")
	}()

	config := tgbotapi.NewUpdate(0)
	config.Timeout = 60

	updates := r.bot.GetUpdatesChan(config)

	for update := range updates {
		r.process(&update)
	}
}
