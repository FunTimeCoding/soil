package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/network"

func (c *Client) DeviceInterfaces(device int32) ([]*network.Interface, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.interfaces[device], nil
}
