package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
)

func (c *Client) DevicePowerPort(
	d *device.Device,
	name string,
) (*power_port.Port, error) {
	result, e := c.DevicePowerPorts(d.Identifier)

	if e != nil {
		return nil, e
	}

	for _, p := range result {
		if p.Name == name {
			return p, nil
		}
	}

	return nil, not_found.Format(
		"power port %s not found for device %s",
		name,
		d.Name,
	)
}
