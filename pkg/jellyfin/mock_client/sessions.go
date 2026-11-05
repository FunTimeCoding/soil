package mock_client

import "github.com/funtimecoding/soil/pkg/jellyfin/session"

func (c *Client) Sessions() ([]*session.Session, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.sessions, nil
}
