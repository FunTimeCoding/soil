package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/cable"
)

func (c *Client) MustDeviceCables(device string) []*cable.Cable {
	result, e := c.DeviceCables(device)
	errors.PanicOnError(e)

	return result
}
