package core

import (
	"openwishlist/app/sdk/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (r *TelegramBot) getTgUserFromUpdate(update *tgbotapi.Update) *models.User {
	if update.Message != nil {
		return &models.User{
			ID: update.Message.From.ID,
		}
	}

	return nil
}
