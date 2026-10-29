package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) SendKey(
	identifier string,
	key string,
) error {
	if _, okay := c.screens[identifier]; !okay {
		return not_found.Format("session %s not found", identifier)
	}

	c.keys[identifier] = append(c.keys[identifier], key)

	return nil
}
