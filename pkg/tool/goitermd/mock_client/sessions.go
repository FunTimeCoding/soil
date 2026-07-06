package mock_client

import "github.com/funtimecoding/soil/pkg/iterm/session"

func (c *Client) Sessions() ([]*session.Session, error) {
	return c.sessions, nil
}
