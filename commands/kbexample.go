package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (ce *CommandExecutor) ExecuteDayword(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
  return
	// If not subscribed:
	//   1. Do you want to subscribe to the word of the day?
	//   2. If yes, confirm the subscription and show today's word
	// If subscribed, show today's word
	// Add a note below that if they wish to unsubscribe, they can do so in settings,
	// typing /settings

	// btns := []utils.KeyboardButton{
	// 	{Name: "Grapes", Inline: true},
	// 	{Name: "Apples", Inline: true},
	// }
	// choice, _ := utils.UserKeyboard("Do you like:", btns, 10, ctx, b, update, additionalContext)
	// b.SendMessage(ctx, &bot.SendMessageParams{
	// 	ChatID: update.Message.Chat.ID,
	// 	Text:   fmt.Sprintf("%s? Meh, sure...", choice),
	// })
}
