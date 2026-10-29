package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/rack"

func (c *Client) RacksByName(_ string) ([]*rack.Rack, error) {
	return nil, c.Failure
}
