package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/module_bay"
)

func (c *Client) MustDeviceModuleBays(device string) []*module_bay.Bay {
	result, e := c.DeviceModuleBays(device)
	errors.PanicOnError(e)

	return result
}
