package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
)

func (c *Client) MustDeviceInterface(
	d *device.Device,
	name string,
) *network.Interface {
	result, e := c.DeviceInterface(d, name)
	errors.PanicOnError(e)

	return result
}
