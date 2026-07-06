package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustDeleteInterfaceByName(
	d *device.Device,
	name string,
) {
	errors.PanicOnError(c.DeleteInterfaceByName(d, name))
}
