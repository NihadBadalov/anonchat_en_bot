package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

  "anonchat_en_bot/utils"
)

func (ce *CommandExecutor) ExecuteFind(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
	*additionalContext = context.WithValue(
		*additionalContext,
		"users_queued",
		append((*additionalContext).Value("users_queued").([]int64), update.Message.Chat.ID),
	)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   `✅️ Searching for a partner...`,
	})

  utils.MatchUsers(ctx, b, additionalContext)
}
