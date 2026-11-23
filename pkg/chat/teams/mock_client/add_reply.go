package mock_client

import "github.com/funtimecoding/soil/pkg/chat/teams/message"

func (c *Client) AddReply(
	team string,
	channel string,
	identifier string,
	m *message.Message,
) {
	k := key(team, channel, identifier)
	c.replies[k] = append(c.replies[k], m)
}
