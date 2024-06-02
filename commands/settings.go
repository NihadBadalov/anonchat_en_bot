package commands

import (
	"anonchat_en_bot/db"
	"anonchat_en_bot/utils"
	"context"
	"fmt"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (ce *CommandExecutor) ExecuteSettings(ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) {
	user, err := db.GetUserByUid(update.Message.Chat.ID)
	if user.Id == 0 || err != nil {
		db.AddUser(update.Message.Chat.ID, -1, -1, true)

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "An error occurred while fetching your settings. Run the command again. If the error persists, try again later.",
		})
		return
	}

	mainMenuButtons := []utils.KeyboardButton{
		{Name: "🚻 Gender 🚻", Inline: false},
		{Name: "🎂 Age 🎂", Inline: false},
		{Name: "📸 Ask permission for photos 📸", Inline: false},
		{Name: "❌ Cancel", Inline: false},
	}
	genderButtons := []utils.KeyboardButton{
		{Name: "👩 Woman", Inline: false},
		{Name: "👨 Man", Inline: false},
		{Name: "❌ Don't disclose", Inline: false},
	}
	photoPermissionButtons := []utils.KeyboardButton{
		{Name: "✅ Yes", Inline: false},
		{Name: "❌ No", Inline: false},
	}

	choice, questionMessage := utils.UserKeyboard("Choose what settings you would like to change", mainMenuButtons, 60, ctx, b, update, additionalContext)
	switch choice {
	case "🚻 Gender 🚻":
		gender, genderQuestionMessage := utils.UserKeyboard("Choose your gender", genderButtons, 60, ctx, b, update, additionalContext, questionMessage)

		var genderInt int
		switch gender {
		case "👩 Woman":
			genderInt = 1
		case "👨 Man":
			genderInt = 0
		case "❌ Don't disclose":
			genderInt = -1
		case "":
			genderInt = -1
		}

		db.SetUserGender(int64(update.Message.ID), genderInt)

		if genderInt != -1 {
      b.DeleteMessage(ctx, &bot.DeleteMessageParams{
        ChatID:   update.Message.Chat.ID,
        MessageID: genderQuestionMessage.ID,
      })
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    update.Message.Chat.ID,
				Text:      fmt.Sprintf("✅ Gender updated to %s", gender),
			})
		}

	case "🎂 Age 🎂":
		age, e := strconv.Atoi(utils.UserInput("Enter your age. You have to be at least 16 years old. Write -1 to cancel.", 60, ctx, b, update, additionalContext))
		if e == nil && age == -1 || (16 <= age && age <= 200) {
			db.SetUserAge(int64(update.Message.ID), age)

      b.DeleteMessage(ctx, &bot.DeleteMessageParams{
        ChatID:   update.Message.Chat.ID,
				MessageID:   questionMessage.ID,
      })
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    update.Message.Chat.ID,
				Text:        fmt.Sprintf("✅ Age updated to %d", age),
			})
		} else {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:      update.Message.Chat.ID,
				Text:        "❌ You have entered an invalid age. Please try again using /settings.",
				ReplyMarkup: nil,
			})

			b.DeleteMessage(ctx, &bot.DeleteMessageParams{
				ChatID:    update.Message.Chat.ID,
				MessageID: questionMessage.ID,
			})
		}

	case "📸 Ask permission for photos 📸":
		choice, questionMessage := utils.UserKeyboard("Do you want to choose whether you want to receive photos every time your partner sends them? If so, press Yes. If you want to receive photos without having to allow, press No.", photoPermissionButtons, 60, ctx, b, update, additionalContext)
		switch choice {
		case "✅ Yes":
			db.SetUserGatekeepMedia(int64(update.Message.ID), true)

			b.EditMessageText(ctx, &bot.EditMessageTextParams{
				ChatID:      update.Message.Chat.ID,
				MessageID:   questionMessage.ID,
				Text:        fmt.Sprintf(""),
				ReplyMarkup: nil,
			})
		case "❌ No":
			db.SetUserGatekeepMedia(int64(update.Message.ID), false)
		}

	case "❌ Cancel":
		b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    update.Message.Chat.ID,
			MessageID: questionMessage.ID,
		})
	}
}
