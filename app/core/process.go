package core

import (
	"context"
	"database/sql"
	"errors"
	"openwishlist/app/sdk"
	"openwishlist/app/sdk/models"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (r *TelegramBot) process(update *tgbotapi.Update) error {
	tgUser := r.getTgUserFromUpdate(update)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := r.dbClient.GetUser(ctx, tgUser); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = r.dbClient.CreateUser(ctx, tgUser)
		}

		if err != nil {
			return err
		}
	}

	switch {
	case update.CallbackQuery != nil:
		return r.handleCallback(tgUser, update)
	case update.Message.Command() != "":
		return r.handleCommand(ctx, tgUser, update)
	}

	return nil
}

func (r *TelegramBot) handleCallback(user *models.User, update *tgbotapi.Update) error {
	return nil
}

func (r *TelegramBot) handleCommand(ctx context.Context, user *models.User, update *tgbotapi.Update) error {
	switch update.Message.Command() {
	case sdk.CommandStart:
		return r.handleStart(user)
	}

	return nil
}
