package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
)

func (c *Client) DeviceInterfaceByName(
	d *device.Device,
	name string,
) (*network.Interface, error) {
	for _, i := range c.interfaces[d.Identifier] {
		if i.Name == name {
			return i, nil
		}
	}

	return nil, not_found.Format(
		"interface %s not found for device %s",
		name,
		d.Name,
	)
}
