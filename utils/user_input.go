package utils

import (
	"context"
	"sync"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func UserInput(s string, cooldown int64, ctx context.Context, b *bot.Bot, update *models.Update, additionalContext *context.Context) string {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   s,
	})

	user_input_timeout := &sync.Map{}
	user_input_timeout.Store("Value", nil)
	user_input_timeout.Store("EndTime", time.Now().Unix()+cooldown)
	(*additionalContext).Value("user_inputs").(*sync.Map).Store(update.Message.Chat.ID, user_input_timeout)

	// WaitGroup to halt the thread until a message appears
	var success bool
	var val string
	var wg sync.WaitGroup

	wg.Add(1)

	go func(ctx *context.Context, sender_id *int64) {
		defer wg.Done()

		cooldown, _ := (*ctx).Value("user_inputs").(*sync.Map).Load(*sender_id)
		var (
			t  any
			_t bool
			v  any
			_v bool
		)

		for {
			// Exceeded EndTime
			t, _t = cooldown.(*sync.Map).Load("EndTime")
			v, _v = cooldown.(*sync.Map).Load("Value")
			if _t && time.Now().Unix() > t.(int64) {
				success = false
				val = ""
				break
			}
			// Got a response
			if _v && v != nil {
				success = true
				val = v.(string)
				break
			}
		}
	}(additionalContext, &update.Message.Chat.ID)

	wg.Wait()

	(*additionalContext).Value("user_inputs").(*sync.Map).Store(update.Message.Chat.ID, nil)

	// User response
	if !success && val == "" {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Response time exceeded",
		})
	}

	return val
}
