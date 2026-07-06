package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
)

func (c *Client) MustDeviceAddresses(device string) []*internet_address.Address {
	result, e := c.DeviceAddresses(device)
	errors.PanicOnError(e)

	return result
}
