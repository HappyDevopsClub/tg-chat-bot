package handlers

import (
	"context"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

func WelcomeMessageFilter(update *models.Update) bool {
	return update.Message != nil && len(update.Message.NewChatMembers) > 0
}

func WelcomeMessageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendSticker(ctx, &bot.SendStickerParams{
		ChatID: update.Message.Chat.ID,
		Sticker: &models.InputFileString{
			Data: os.Getenv("WELCOME_STICKER_ID"),
		},
	})
	if err != nil {
		log.Error("Failed to send welcome message", zap.Error(err))
	}
}
