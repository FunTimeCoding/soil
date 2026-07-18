package telegram

import (
	"github.com/funtimecoding/soil/pkg/chat/telegram/client"
	"github.com/funtimecoding/soil/pkg/chat/telegram/store"
)

func New(
	token string,
	s *store.Store,
) *Client {
	// https://github.com/go-telegram-bot-api/telegram-bot-api
	return &Client{client: client.New(token), store: s}
}
