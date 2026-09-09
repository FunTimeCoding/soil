package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateInterface(
	_ *device.Device,
	_ string,
	_ netbox.InterfaceTypeValue,
) (*network.Interface, error) {
	return nil, nil
}
