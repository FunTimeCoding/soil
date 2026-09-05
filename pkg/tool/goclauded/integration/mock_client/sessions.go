package mock_client

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"

func (c *Client) Sessions() []*session.Session {
	var result []*session.Session

	for _, s := range c.sessions {
		result = append(result, s)
	}

	return result
}
