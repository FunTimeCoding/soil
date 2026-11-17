package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
)

func (c *Client) DevicePowerOutlet(
	d *device.Device,
	name string,
) (*power_outlet.Outlet, error) {
	result, e := c.DevicePowerOutlets(d.Identifier)

	if e != nil {
		return nil, e
	}

	for _, o := range result {
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
