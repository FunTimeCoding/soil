package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/location"

func (c *Client) CreateLocation(
	_ string,
	_ string,
) (*location.Location, error) {
	return nil, nil
}
