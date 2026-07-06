package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
)

func (c *Client) DeviceInterfaceByName(
	d *device.Device,
	name string,
) (*network.Interface, error) {
	result, e := c.DeviceInterfaces(d.Identifier)

	if e != nil {
		return nil, e
	}

	for _, i := range result {
		if i.Name == name {
			return i, nil
		}
	}

	return nil, fmt.Errorf(
		"interface %s not found for device %s: %w",
		name,
		d.Name,
		constant.ErrorNotFound,
	)
}
