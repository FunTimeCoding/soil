package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) SetTabColor(
	sessionIdentifier string,
	_ int,
	_ int,
	_ int,
) error {
	if _, okay := c.screens[sessionIdentifier]; !okay {
		return not_found.Format("session %s not found", sessionIdentifier)
	}

	return nil
}
