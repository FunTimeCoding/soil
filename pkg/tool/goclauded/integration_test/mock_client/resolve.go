package mock_client

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"

func (c *Client) Resolve(query string) *session.Session {
	if s, okay := c.sessions[query]; okay {
		return s
	}

	return session.Stub()
}
