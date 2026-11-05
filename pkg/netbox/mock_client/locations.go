package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/location"

func (c *Client) Locations() ([]*location.Location, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
