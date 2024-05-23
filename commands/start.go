package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (ce *CommandExecutor) ExecuteStart(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: `Welcome to Anonymous Chat! 🎉️
Here are some commands you can use:`,
	})

  ce.ExecuteHelp(ctx, b, update, additionalContext)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: `_Alternatively, use /help to show this message again._`,
	})
}
