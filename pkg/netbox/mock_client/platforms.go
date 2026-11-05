package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/platform"

func (c *Client) Platforms() ([]*platform.Platform, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
