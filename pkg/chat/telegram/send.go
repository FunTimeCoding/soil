package telegram

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Client) Send(m tgbotapi.Chattable) {
	_, e := c.client.Send(m)
	errors.PanicOnError(e)
}
