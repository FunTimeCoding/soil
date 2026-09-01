package discord

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) DeleteLoop(channel string) {
	for {
		messages, e := c.client.ChannelMessages(
			channel,
			constant.DiscordMessageLimit,
			"",
			"",
			"",
		)
		errors.PanicOnError(e)

		if len(messages) == 0 {
			break
		}

		for i, m := range messages {
			console.Format("Delete %d: %s\n", i, m.ID)
			c.Delete(channel, m)
		}
	}
}
