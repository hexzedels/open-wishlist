package core

import (
	"openwishlist/app/sdk"
	"openwishlist/app/sdk/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (r *TelegramBot) getTgUserFromUpdate(update *tgbotapi.Update) *models.User {
	return &models.User{
		ID:       update.SentFrom().ID,
		Username: update.SentFrom().UserName,
		State:    sdk.StateHome,
	}
}
