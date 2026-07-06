package telegram

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Client) Updates() []tgbotapi.Update {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	result, e := c.client.GetUpdates(u)
	errors.PanicOnError(e)

	return result
}
