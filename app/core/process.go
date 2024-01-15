package core

import (
	"openwishlist/app/sdk/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (r *TelegramBot) process(update *tgbotapi.Update) error {
	tgUser := r.getTgUserFromUpdate(update)

	if err := r.dbClient.GetUser(tgUser); err != nil {
		return err
	}

	switch {
	case update.CallbackQuery != nil:
		return r.handleCallback(tgUser, update)
	}

	return nil
}

func (r *TelegramBot) handleCallback(user *models.User, update *tgbotapi.Update) error {
	return nil
}
