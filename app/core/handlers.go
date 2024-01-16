package core

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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
				sdk.ButtonNewItem,
			),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(
				sdk.ButtonListItems,
			),
		),
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

	config := tgbotapi.NewMessage(user.ID, fmt.Sprintf(sdk.MessageWishlistCreated, update.Message.Text))
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

func (r *TelegramBot) handleChooseWishlist(ctx context.Context, user *models.User, update *tgbotapi.Update) error {
	wishlists, err := r.dbClient.ListWishlists(ctx, user)
	if err != nil {
		return err
	}

	var inlineMarkup [][]tgbotapi.InlineKeyboardButton

	for _, list := range wishlists {
		inlineMarkup = append(
			inlineMarkup,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					list.Name,
					sdk.WrapCallback(
						sdk.CallbackWishlist,
						strconv.Itoa(int(list.ID)),
					),
				)))
	}

	config := tgbotapi.NewMessage(user.ID, sdk.MessageWishlistList)
	config.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: inlineMarkup}

	if _, err := r.bot.Send(config); err != nil {
		user.State = sdk.StateHome
		if err := r.dbClient.UpdateUser(ctx, user); err != nil {
			r.logger.Error("failed to update user after failed bot answer", zap.Error(err))
		}

		return err
	}

	user.State = sdk.StateWishlistChoose

	return r.dbClient.UpdateUser(ctx, user)
}

func (r *TelegramBot) handleWishlistSet(ctx context.Context, user *models.User, id string) error {
	wid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	user.WishlistID = int64(wid)
	user.State = sdk.StateHome

	wishlist := &models.Wishlist{
		ID: user.WishlistID,
	}

	r.dbClient.GetWishlist(ctx, wishlist)
	config := tgbotapi.NewMessage(user.ID, fmt.Sprintf(sdk.MessageWishlistSet, wishlist.Name))

	if _, err := r.bot.Send(config); err != nil {
		user.State = sdk.StateHome
		if err := r.dbClient.UpdateUser(ctx, user); err != nil {
			r.logger.Error("failed to update user after failed bot answer", zap.Error(err))
		}

		return err
	}

	return r.dbClient.UpdateUser(ctx, user)
}

func (r *TelegramBot) handleNewItemButton(ctx context.Context, user *models.User) error {
	user.State = sdk.StateItemNew

	config := tgbotapi.NewMessage(user.ID, sdk.MessageItemNew)
	if _, err := r.bot.Send(config); err != nil {
		user.State = sdk.StateHome
		if err := r.dbClient.UpdateUser(ctx, user); err != nil {
			r.logger.Error("failed to update user after failed bot answer", zap.Error(err))
		}

		return err
	}

	return r.dbClient.UpdateUser(ctx, user)
}

func (r *TelegramBot) handleItemName(ctx context.Context, user *models.User, update *tgbotapi.Update) error {
	user.State = sdk.StateHome

	item := &models.Item{
		WishlistID: user.WishlistID,
		Name:       update.Message.Text,
	}

	r.dbClient.CreateItem(ctx, item)

	config := tgbotapi.NewMessage(user.ID, sdk.MessageCreatedItem)
	if _, err := r.bot.Send(config); err != nil {
		user.State = sdk.StateHome
		if err := r.dbClient.UpdateUser(ctx, user); err != nil {
			r.logger.Error("failed to update user after failed bot answer", zap.Error(err))
		}

		return err
	}

	return r.dbClient.UpdateUser(ctx, user)
}

func (r *TelegramBot) handleButtonListItems(ctx context.Context, user *models.User) error {
	items, err := r.dbClient.ListItems(ctx, user)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(sdk.MessagedListItems)
	for i, item := range items {
		b.WriteString(fmt.Sprintf(sdk.MessageTableFormat, i, item.ID, item.Name))
	}

	config := tgbotapi.NewMessage(user.ID, b.String())
	config.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := r.bot.Send(config); err != nil {
		user.State = sdk.StateHome
		if err := r.dbClient.UpdateUser(ctx, user); err != nil {
			r.logger.Error("failed to update user after failed bot answer", zap.Error(err))
		}

		return err
	}

	return nil
}
