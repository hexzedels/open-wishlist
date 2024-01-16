package core

import (
	"context"
	"strconv"

	"openwishlist/app/sdk"
	"openwishlist/app/sdk/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (r *TelegramBot) handleStart(user *models.User) error {
	config := tgbotapi.NewMessage(user.ID, sdk.MessageStart)
	config.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(
				sdk.ButtonNewWishlist,
			),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(
				sdk.ButtonExistingWishlist,
			),
		),
	)

	if _, err := r.bot.Send(config); err != nil {
		return err
	}

	return nil
}

func (r *TelegramBot) handleNewWishlistButton(ctx context.Context, user *models.User) error {
	user.State = sdk.StateWishlistNew
	if err := r.dbClient.UpdateUser(ctx, user); err != nil {
		return err
	}

	config := tgbotapi.NewMessage(user.ID, sdk.MessageWishlistNew)
	if _, err := r.bot.Send(config); err != nil {
		user.State = sdk.StateHome
		if err := r.dbClient.UpdateUser(ctx, user); err != nil {
			r.logger.Error("failed to update user after failed bot answer", zap.Error(err))
		}

		return err
	}

	return nil
}

func (r *TelegramBot) handleNewWishlist(ctx context.Context, user *models.User, update *tgbotapi.Update) error {
	if err := r.dbClient.CreateWishlist(
		ctx,
		&models.Wishlist{
			OwnerID: user.ID,
			Name:    update.Message.Text,
		},
	); err != nil {
		return err
	}

	user.State = sdk.StateHome

	return r.dbClient.UpdateUser(ctx, user)
}

func (r *TelegramBot) handleListAllWishlists(ctx context.Context, user *models.User, update *tgbotapi.Update) error {
	wishlists, err := r.dbClient.ListWishlists(ctx, user)
	if err != nil {
		return err
	}

	var inlineMarkup []tgbotapi.InlineKeyboardButton

	for _, list := range wishlists {
		inlineMarkup = append(
			inlineMarkup,
			tgbotapi.NewInlineKeyboardButtonData(
				list.Name,
				sdk.WrapCallback(
					sdk.CallbackWishlist,
					strconv.Itoa(int(list.ID)),
				),
			))
	}

	config := tgbotapi.NewMessage(user.ID, sdk.MessageWishlistList)
	config.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(inlineMarkup)

	if _, err := r.bot.Send(config); err != nil {
		user.State = sdk.StateHome
		if err := r.dbClient.UpdateUser(ctx, user); err != nil {
			r.logger.Error("failed to update user after failed bot answer", zap.Error(err))
		}

		return err
	}

	user.State = sdk.StateHome

	return r.dbClient.UpdateUser(ctx, user)
}
