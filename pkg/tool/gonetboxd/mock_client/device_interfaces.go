package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/network"

func (c *Client) DeviceInterfaces(_ int32) ([]*network.Interface, error) {
	return nil, nil
}
