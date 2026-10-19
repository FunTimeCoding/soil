package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) SendKey(
	identifier string,
	_ string,
) error {
	if _, okay := c.screens[identifier]; !okay {
		return not_found.Format("session %s not found", identifier)
	}

	return nil
}
