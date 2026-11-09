package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
)

func (c *Client) MustDeviceType(n string) *device_type.Type {
	result, e := c.DeviceType(n)
	errors.PanicOnError(e)

	return result
}
