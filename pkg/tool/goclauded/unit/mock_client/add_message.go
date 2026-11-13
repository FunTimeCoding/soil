package mock_client

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"

func (c *Client) AddMessage(
	sessionIdentifier string,
	m *message.Message,
) {
	c.messages[sessionIdentifier] = append(c.messages[sessionIdentifier], *m)
}
