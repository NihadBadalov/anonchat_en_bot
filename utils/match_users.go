package utils

import (
	"context"
	"sync"
	"time"

	"github.com/go-telegram/bot"
)

const matchMessage = `🎉 You have been matched with a partner!

/next - find a new partner
/stop - stop this chat

@AnonnCBot`

func MatchUsers(
	ctx context.Context,
	b *bot.Bot,
	additionalContext *context.Context,
) {
	usersQueued, ok := (*additionalContext).Value("users_queued").([]int64)
	if !ok {
		usersQueued = []int64{}
	}

	// If there are no users in the queue, return
	if len(usersQueued) < 2 {
		return
	}

	// Get the first two users from the queue
	user1 := usersQueued[0]
	user2 := usersQueued[1]

	// Remove the two users from the queue
	usersQueued = usersQueued[2:]
	*additionalContext = context.WithValue(
		*additionalContext,
		"users_queued",
		usersQueued,
	)

	// Send a message to the two users
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: user1,
		Text:   matchMessage,
	})
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: user2,
		Text:   matchMessage,
	})

	// Add the two users to the matched users list
	t := time.Now().Unix()
	chat := &sync.Map{}
	chat.Store("u1", user1)
	chat.Store("u2", user2)
	chat.Store("time", t)
	(*additionalContext).Value("user_chats").(*sync.Map).Store(user1, chat)
	(*additionalContext).Value("user_chats").(*sync.Map).Store(user2, chat)
}
