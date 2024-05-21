package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandExecutor struct{}

func (ce *CommandExecutor) ExecuteHelp(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: `- anonchat_en_bot

/find - to find a new partner
/preferences - to set your preferences
/cancel - to cancel the search`,
	})
}
