package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/network"
)

func (c *Client) MustDeviceInterfaceByName(
	d *device.Device,
	name string,
) *network.Interface {
	result, e := c.DeviceInterfaceByName(d, name)
	errors.PanicOnError(e)

	return result
}
