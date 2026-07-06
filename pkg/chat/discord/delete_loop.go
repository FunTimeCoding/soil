package discord

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) DeleteLoop(channel string) {
	for {
		messages, e := c.client.ChannelMessages(
			channel,
			MessageLimit,
			"",
			"",
			"",
		)
		errors.PanicOnError(e)

		if len(messages) == 0 {
			break
		}

		for i, m := range messages {
			fmt.Printf("Delete %d: %s\n", i, m.ID)
			c.Delete(channel, m)
		}
	}
}
