package commands

import (
	"context"
	"sync"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"anonchat_en_bot/utils"
)

func removeFromSlice(slice []int64, id int64) ([]int64, bool) {
	for i, v := range slice {
		if v == id {
			return append(slice[:i], slice[i+1:]...), true
		}
	}
	return slice, false
}

func (ce *CommandExecutor) ExecuteCancel(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
	newUsersQueued, found := removeFromSlice((*additionalContext).Value("users_queued").([]int64), update.Message.Chat.ID)

	if found {
		// The user was in the queue
		*additionalContext = context.WithValue(
			*additionalContext,
			"users_queued",
			newUsersQueued,
		)

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   `✅️ Search canceled.`,
		})

		utils.MatchUsers(ctx, b, additionalContext)
	} else {
		if chatObj, ok := (*additionalContext).Value("user_chats").(*sync.Map).Load(update.Message.Chat.ID); ok && chatObj != nil {
			u1, ok1 := chatObj.(*sync.Map).Load("u1")
			u2, ok2 := chatObj.(*sync.Map).Load("u2")
			var opposingUser int64
			if ok1 && u1 != nil && u1.(int64) == update.Message.Chat.ID && ok2 && u2 != nil {
				opposingUser = u2.(int64)
			} else {
				opposingUser = u1.(int64)
			}

			(*additionalContext).Value("user_chats").(*sync.Map).Delete(update.Message.Chat.ID)
			(*additionalContext).Value("user_chats").(*sync.Map).Delete(opposingUser)

			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   `✅️ Chat canceled.`,
			})

			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: opposingUser,
				Text:   `❌️ Chat canceled by partner.`,
			})
		}
	}
}
