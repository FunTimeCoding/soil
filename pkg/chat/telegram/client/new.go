package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func New(token string) *tgbotapi.BotAPI {
	result, e := tgbotapi.NewBotAPI(token)
	errors.PanicOnError(e)

	return result
}
