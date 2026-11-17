package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetDeviceSite(
	name string,
	siteName string,
) (*device.Device, error) {
	d, e := c.Device(name)

	if e != nil {
		return nil, e
	}

	s, f := c.Site(siteName)

	if f != nil {
		return nil, f
	}

	q := netbox.NewPatchedWritableDeviceWithConfigContextRequest()
	q.SetSite(netbox.DeviceWithConfigContextRequestSite{Int32: &s.Identifier})

	return c.updateDevice(d, q)
}
