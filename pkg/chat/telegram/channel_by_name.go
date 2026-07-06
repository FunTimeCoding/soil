package telegram

import "github.com/funtimecoding/soil/pkg/chat/telegram/database/channel"

func (c *Client) ChannelByName(name string) *channel.Channel {
	for _, h := range c.channels {
		if h.Name == name {
			return h
		}
	}

	return nil
}
