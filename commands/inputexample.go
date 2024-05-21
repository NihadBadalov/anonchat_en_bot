package commands

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"anonchat_en_bot/utils"
)

func (ce *CommandExecutor) ExecuteInputexample(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
	user_input := utils.UserInput(
		"Whatcha name, bruv?! Got 10 seconds to ansa",
		10,
		ctx,
		b,
		update,
		additionalContext,
	)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("%s? Noice name, bruv, innit?", user_input),
	})
}
