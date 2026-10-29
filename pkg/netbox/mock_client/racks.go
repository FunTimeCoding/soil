package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/rack"

func (c *Client) Racks() ([]*rack.Rack, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return c.racks, nil
}
