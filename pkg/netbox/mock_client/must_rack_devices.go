package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustRackDevices(i int32) []*device.Device {
	result, e := c.RackDevices(i)
	errors.PanicOnError(e)

	return result
}
