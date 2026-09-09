package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
)

func (c *Client) DeviceInterfaceByName(
	_ *device.Device,
	_ string,
) (*network.Interface, error) {
	return nil, nil
}
