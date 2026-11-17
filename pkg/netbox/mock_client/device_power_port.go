package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
)

func (c *Client) DevicePowerPort(
	d *device.Device,
	name string,
) (*power_port.Port, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, p := range c.ports[d.Identifier] {
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
