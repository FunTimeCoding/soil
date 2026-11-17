package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
)

func (c *Client) DevicePowerOutlet(
	d *device.Device,
	name string,
) (*power_outlet.Outlet, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, o := range c.outlets[d.Identifier] {
		if o.Name == name {
			return o, nil
		}
	}

	return nil, not_found.Format(
		"power outlet %s not found for device %s",
		name,
		d.Name,
	)
}
