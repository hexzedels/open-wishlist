package core

import (
	"context"
	"database/sql"
	"errors"
	"openwishlist/app/sdk"
	"openwishlist/app/sdk/models"
	"strconv"
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
		return r.handleCallback(ctx, tgUser, update)
	case update.Message.Command() != "":
		return r.handleCommand(ctx, tgUser, update)
	case update.Message.Text != "":
		return r.handleMessage(ctx, tgUser, update)
	}

	return nil
}

func (r *TelegramBot) handleCallback(ctx context.Context, user *models.User, update *tgbotapi.Update) error {
	call, id := sdk.UnwrapCallback(update.CallbackQuery.Data)

	switch call {
	case sdk.CallbackWishlist:
		return r.handleWishlistSet(ctx, user, id)
	}
	return nil
}

func (r *TelegramBot) handleWishlistSet(ctx context.Context, user *models.User, id string) error {
	wid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	user.ID = int64(wid)

	return r.dbClient.UpdateUser(ctx, user)
}

func (r *TelegramBot) handleCommand(ctx context.Context, user *models.User, update *tgbotapi.Update) error {
	switch update.Message.Command() {
	case sdk.CommandStart:
		return r.handleStart(user)
	}

	return nil
}

func (r *TelegramBot) handleMessage(ctx context.Context, user *models.User, update *tgbotapi.Update) error {
	switch update.Message.Text {
	case sdk.ButtonNewWishlist:
		return r.handleNewWishlistButton(ctx, user)
	case sdk.ButtonExistingWishlist:
		return r.handleListAllWishlists(ctx, user, update)
	}

	switch user.State {
	case sdk.StateWishlistNew:
		return r.handleNewWishlist(ctx, user, update)
	}

	return nil
}
