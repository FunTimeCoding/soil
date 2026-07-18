package telegram

import "github.com/funtimecoding/soil/pkg/chat/telegram/database/channel"

func (c *Client) ChannelByName(name string) *channel.Channel {
	if c.store == nil {
		return nil
	}

	return c.store.MustChannelByName(name)
}
