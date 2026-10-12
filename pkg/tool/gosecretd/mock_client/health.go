package mock_client

import "github.com/funtimecoding/soil/pkg/openbao/status"

func (c *Client) Health() (*status.Status, error) {
	if c.health != nil {
		return nil, c.health
	}

	return c.status, nil
}
