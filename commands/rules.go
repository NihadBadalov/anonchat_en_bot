package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (ce *CommandExecutor) ExecuteRules(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: `- anonchat_en_bot

- Do not spam.
- Do not send NSFW or pornographic content.
- Do not send personal information.
- Do not send malicious links.
- Do not advertise other bots or channels.
- Do not send offensive content.`,
	})
}
