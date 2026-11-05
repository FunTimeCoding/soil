package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/cable"

func (c *Client) Cables() ([]*cable.Cable, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
