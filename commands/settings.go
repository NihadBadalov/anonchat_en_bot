package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (ce *CommandExecutor) ExecuteSettings(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
  ce.ExecutePreferences(ctx, b, update, additionalContext)
}
