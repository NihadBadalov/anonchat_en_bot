package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

  "anonchat_en_bot/utils"
)

func removeFromSlice(slice []int64, id int64) []int64 {
  for i, v := range slice {
    if v == id {
      return append(slice[:i], slice[i+1:]...)
    }
  }
  return slice
}

func (ce *CommandExecutor) ExecuteCancel(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
  *additionalContext = context.WithValue(
    *additionalContext,
    "users_queued",
    removeFromSlice((*additionalContext).Value("users_queued").([]int64), update.Message.Chat.ID),
  )

  b.SendMessage(ctx, &bot.SendMessageParams{
    ChatID: update.Message.Chat.ID,
    Text:   `✅️ Search canceled.`,
  })

  utils.MatchUsers(ctx, b, additionalContext)
}
