package mock_client

import "github.com/funtimecoding/soil/pkg/jellyfin/session"

func (c *Client) AddSession(s *session.Session) {
	c.sessions = append(c.sessions, s)
}
