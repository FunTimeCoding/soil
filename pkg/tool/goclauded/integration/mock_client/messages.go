package mock_client

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"

func (c *Client) Messages(sessionIdentifier string) []message.Message {
	return nil
}
