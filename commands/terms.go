package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (ce *CommandExecutor) ExecuteTerms(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: `- @AnonnCBot

- For now, NO information AT ALL is collected.
- Please, check Terms regurarly for updates.

Last updated: 2024-05-23`,
	})
}
