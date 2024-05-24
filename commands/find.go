package commands

import (
	"context"
	"sync"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"anonchat_en_bot/utils"
)

func (ce *CommandExecutor) ExecuteFind(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
	if val, ok := (*additionalContext).Value("user_chats").(*sync.Map).Load(update.Message.Chat.ID); ok && val != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   `❌️ You are already in a chat. Use /cancel to cancel it.`,
		})

		utils.MatchUsers(ctx, b, additionalContext)
		return
	}

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
