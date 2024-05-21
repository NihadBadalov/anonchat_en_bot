package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"sync"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"anonchat_en_bot/commands"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	additionalContext, cancelInput := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancelInput()
	defer cancel()

	ce := &commands.CommandExecutor{}
	// user_inputs is id->unix
	additionalContext = context.WithValue(additionalContext, "user_inputs", &sync.Map{})
	additionalContext = context.WithValue(additionalContext, "user_keyboards", &sync.Map{})
	additionalContext = context.WithValue(additionalContext, "user_chats", &sync.Map{})
	additionalContext = context.WithValue(additionalContext, "users_queued", []int64{})

	opts := []bot.Option{
		bot.WithDefaultHandler(handler(&additionalContext)),
		bot.WithCallbackQueryDataHandler("btn_", bot.MatchTypePrefix, keyboardHandler(&additionalContext)),
	}

	b, err := bot.New(os.Getenv("BOT_TOKEN"), opts...)
	if err != nil {
		panic(err)
	}

	b.RegisterHandlerMatchFunc(isCommand(b, ctx, ce), commandHandler(ce, &additionalContext))

	fmt.Println("Bot started")
	b.Start(ctx)
	fmt.Println("Bot stopped")
}

func isCommand(b *bot.Bot, ctx context.Context, ce *commands.CommandExecutor) func(*models.Update) bool {
	return func(update *models.Update) bool {
		if len([]rune(update.Message.Text)) < 2 {
			return false
		}
    if update.Message.Text[0] != '/' {
      return false
    }

		cmdName := strings.Split(update.Message.Text, " ")[0]
		executeFunctionName := "Execute" + strings.ToUpper(string(rune(cmdName[1]))) + cmdName[2:]
		cmdValue := reflect.ValueOf(ce).MethodByName(executeFunctionName)
		if !cmdValue.IsValid() {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Command not found! oopsie-daisy...",
			})
			return false
		}
		return true
	}
}

func commandHandler(ce *commands.CommandExecutor, additionalContext *context.Context) func(context.Context, *bot.Bot, *models.Update) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message.Text[0] != '/' {
			return
		}

		cmdName := strings.Split(update.Message.Text, " ")[0]
		executeFunctionName := "Execute" + strings.ToUpper(string(rune(cmdName[1]))) + cmdName[2:]
		cmdValue := reflect.ValueOf(ce).MethodByName(executeFunctionName)
		if !cmdValue.IsValid() {
			return
		}

		args := []reflect.Value{
			reflect.ValueOf(ctx),
			reflect.ValueOf(b),
			reflect.ValueOf(update),
			reflect.ValueOf(additionalContext),
		}

		// Using "go" to run the function in a goroutine
		go cmdValue.Call(args)
	}
}

func keyboardHandler(additionalContext *context.Context) func(ctx context.Context, b *bot.Bot, update *models.Update) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})

		// Check if we are waiting for user input
		if val, ok := (*additionalContext).Value("user_keyboards").(*sync.Map).Load(update.CallbackQuery.From.ID); ok && val != nil {
			v, success := (*additionalContext).Value("user_keyboards").(*sync.Map).Load(update.CallbackQuery.From.ID)
			if success {
				v.(*sync.Map).Store("Value", update.CallbackQuery.Data)
			}
		}
	}
}

func handler(additionalContext *context.Context) func(context.Context, *bot.Bot, *models.Update) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		// Handler for other messages that are NOT commands
		if val, ok := (*additionalContext).Value("user_inputs").(*sync.Map).Load(update.Message.Chat.ID); ok && val != nil {
			v, success := (*additionalContext).Value("user_inputs").(*sync.Map).Load(update.Message.Chat.ID)
			if success {
				v.(*sync.Map).Store("Value", update.Message.Text)
				return
			}
		}

		// If the user is talking to a partner
		if val, ok := (*additionalContext).Value("user_chats").(*sync.Map).Load(update.Message.Chat.ID); ok && val != nil {
			u1, ok1 := val.(*sync.Map).Load("u1")
			u2, ok2 := val.(*sync.Map).Load("u2")
			var opposingUser int64
			if ok1 && u1 != nil && u1.(int64) == update.Message.Chat.ID && ok2 && u2 != nil {
				opposingUser = u2.(int64)
			} else {
				opposingUser = u1.(int64)
			}

			// Handle sending the message
			// It's a sticker
			if update.Message.Sticker != nil {
				b.SendSticker(ctx, &bot.SendStickerParams{
					ChatID: opposingUser,
					Sticker: &models.InputFileString{
						Data: update.Message.Sticker.FileID,
					},
				})
				return
			}
			if update.Message.Photo != nil {
        var medias []models.InputMedia
        for _, photo := range update.Message.Photo {
          medias = append(medias, &models.InputMediaPhoto{
            Media: photo.FileID,
            HasSpoiler: true,
            Caption: update.Message.Caption,
            CaptionEntities: update.Message.CaptionEntities,
          })
        }

        b.SendMediaGroup(ctx, &bot.SendMediaGroupParams{
          ChatID: opposingUser,
          Media: medias,
        })
				return
			}
			if update.Message.Video != nil {
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: update.Message.Chat.ID,
					Text:   "Videos are not supported",
				})
				return
			}
			if update.Message.Voice != nil {
				b.SendVoice(ctx, &bot.SendVoiceParams{
					ChatID: opposingUser,
					Voice: &models.InputFileString{
						Data: update.Message.Voice.FileID,
					},
				})
				return
			}
			if update.Message.Animation != nil {
				b.SendAnimation(ctx, &bot.SendAnimationParams{
					ChatID: opposingUser,
					Animation: &models.InputFileString{
						Data: update.Message.Animation.FileID,
					},
				})
			}
			if update.Message.Text != "" {
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: opposingUser,
					Text:   update.Message.Text,
				})
				return
			}
		}
	}
}
