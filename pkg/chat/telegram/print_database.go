package telegram

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func (c *Client) PrintDatabase() {
	if c.store == nil {
		return
	}

	f := constant.TelegramFormat
	console.Line("Channels:")

	for _, h := range c.store.MustChannels() {
		console.Line(h.Format(f))
	}

	console.Line()
	console.Line("Users:")

	for _, u := range c.store.MustUsers() {
		console.Line(u.Format(f))
	}
}
