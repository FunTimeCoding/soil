package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/location"

func (c *Client) Location(_ string) (*location.Location, error) {
	return nil, c.Failure
}
