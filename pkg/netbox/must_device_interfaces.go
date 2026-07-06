package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/network"
)

func (c *Client) MustDeviceInterfaces(device int32) []*network.Interface {
	result, e := c.DeviceInterfaces(device)
	errors.PanicOnError(e)

	return result
}
