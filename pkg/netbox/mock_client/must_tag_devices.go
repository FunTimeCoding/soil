package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustTagDevices(s string) []*device.Device {
	result, e := c.TagDevices(s)
	errors.PanicOnError(e)

	return result
}
