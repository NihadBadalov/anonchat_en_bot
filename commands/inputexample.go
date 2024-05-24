package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (ce *CommandExecutor) ExecuteInputexample(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
  return
	// user_input := utils.UserInput(
	// 	"Whatcha name, bruv?! Got 10 seconds to ansa",
	// 	10,
	// 	ctx,
	// 	b,
	// 	update,
	// 	additionalContext,
	// )
	//
	// b.SendMessage(ctx, &bot.SendMessageParams{
	// 	ChatID: update.Message.Chat.ID,
	// 	Text:   fmt.Sprintf("%s? Noice name, bruv, innit?", user_input),
	// })
}
