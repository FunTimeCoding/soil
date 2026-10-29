package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/location"

func (c *Client) LocationByName(_ string) (*location.Location, error) {
	return nil, c.Failure
}
