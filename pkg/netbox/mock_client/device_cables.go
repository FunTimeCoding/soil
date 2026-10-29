package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/cable"

func (c *Client) DeviceCables(_ string) ([]*cable.Cable, error) {
	return nil, c.Failure
}
