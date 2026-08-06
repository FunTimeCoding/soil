package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetDevicePrimaryAddress(
	name string,
	address string,
) (*device.Device, error) {
	d, e := c.DeviceByName(name)

	if e != nil {
		return nil, e
	}

	addresses, f := c.DeviceAddresses(name)

	if f != nil {
		return nil, f
	}

	match := internet_address.Find(addresses, address)

	if match == nil {
		return nil, fmt.Errorf(
			"address %s not assigned to device %s",
			address,
			name,
		)
	}

	q := netbox.NewPatchedWritableDeviceWithConfigContextRequest()
	q.SetPrimaryIp4(
		netbox.DeviceWithConfigContextRequestPrimaryIp4{
			Int32: &match.Identifier,
		},
	)

	return c.updateDevice(d, q)
}
