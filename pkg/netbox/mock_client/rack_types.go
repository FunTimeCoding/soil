package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/rack_type"

func (c *Client) RackTypes() ([]*rack_type.Type, error) {
	return nil, c.Failure
}
