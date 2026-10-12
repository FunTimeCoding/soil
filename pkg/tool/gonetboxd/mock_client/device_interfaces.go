package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/network"

func (c *Client) DeviceInterfaces(device int32) ([]*network.Interface, error) {
	return c.interfaces[device], nil
}
