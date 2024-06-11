package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (ce *CommandExecutor) ExecuteSearch(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
  ce.ExecuteFind(ctx, b, update, additionalContext)
}
