package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) updateDevice(
	d *device.Device,
	q *netbox.PatchedWritableDeviceWithConfigContextRequest,
) (*device.Device, error) {
	result, _, e := c.client.DcimAPI.DcimDevicesPartialUpdate(
		c.context,
		d.Identifier,
	).PatchedWritableDeviceWithConfigContextRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return device.New(result), nil
}
