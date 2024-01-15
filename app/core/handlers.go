package core

import (
	"openwishlist/app/sdk"
	"openwishlist/app/sdk/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (r *TelegramBot) handleStart(user *models.User) error {
	if _, err := r.bot.Send(tgbotapi.NewMessage(user.ID, sdk.MessageStart)); err != nil {
		return err
	}

	return nil
}
