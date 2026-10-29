package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/rack"

func (c *Client) RackByName(_ string) (*rack.Rack, error) {
	return nil, c.Failure
}
