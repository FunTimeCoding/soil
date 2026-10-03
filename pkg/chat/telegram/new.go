package telegram

import (
	"github.com/funtimecoding/soil/pkg/chat/telegram/client"
	"github.com/funtimecoding/soil/pkg/chat/telegram/store"
)

// Reference: https://github.com/go-telegram-bot-api/telegram-bot-api
func New(
	token string,
	s *store.Store,
) *Client {
	return &Client{client: client.New(token), store: s}
}
