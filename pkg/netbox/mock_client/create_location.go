package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/location"

func (c *Client) CreateLocation(
	_ string,
	_ string,
) (*location.Location, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
