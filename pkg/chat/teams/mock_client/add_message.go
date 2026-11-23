package mock_client

import "github.com/funtimecoding/soil/pkg/chat/teams/message"

func (c *Client) AddMessage(
	team string,
	channel string,
	m *message.Message,
) {
	k := key(team, channel)
	c.messages[k] = append(c.messages[k], m)
}
