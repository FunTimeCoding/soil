package telegram

import (
	"github.com/funtimecoding/soil/pkg/chat/telegram/store"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	client  *tgbotapi.BotAPI
	Verbose bool
	store   *store.Store
}
