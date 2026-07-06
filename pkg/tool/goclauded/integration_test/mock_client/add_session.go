package mock_client

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"

func (c *Client) AddSession(s *session.Session) {
	c.sessions[s.Identifier] = s
}
