package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateInterface(
	d *device.Device,
	name string,
	t netbox.InterfaceTypeValue,
) (*network.Interface, error) {
	result := c.SeedInterface(d, name)
	result.Type = t

	return result, nil
}
