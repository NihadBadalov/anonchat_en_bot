package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (ce *CommandExecutor) ExecuteNext(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
  successfullCancel := ce.ExecuteCancel(ctx, b, update, additionalContext)

  if !successfullCancel {
    ce.ExecuteFind(ctx, b, update, additionalContext)
  }
}
