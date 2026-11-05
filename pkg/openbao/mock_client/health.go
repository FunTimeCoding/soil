package mock_client

import "github.com/funtimecoding/soil/pkg/openbao/status"

func (c *Client) Health() (*status.Status, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if c.health != nil {
		return nil, c.health
	}

	return c.status, nil
}
