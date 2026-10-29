package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (c *Client) UpdateInterface(
	_ *device.Device,
	_ string,
	_ netbox.InterfaceTypeValue,
	_ net.HardwareAddr,
) (*network.Interface, error) {
	return nil, nil
}
