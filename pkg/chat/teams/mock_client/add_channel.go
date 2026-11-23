package mock_client

import "github.com/funtimecoding/soil/pkg/chat/teams/channel"

func (c *Client) AddChannel(
	team string,
	h *channel.Channel,
) {
	c.channels[team] = append(c.channels[team], h)
}
