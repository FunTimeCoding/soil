package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetDeviceTenant(
	name string,
	tenantName string,
) (*device.Device, error) {
	d, e := c.DeviceByName(name)

	if e != nil {
		return nil, e
	}

	n, f := c.TenantByName(tenantName)

	if f != nil {
		return nil, f
	}

	q := netbox.NewPatchedWritableDeviceWithConfigContextRequest()
	q.SetTenant(
		netbox.BriefTenantRequestAsASNRangeRequestTenant(
			netbox.NewBriefTenantRequest(n.Name, n.Raw.Slug),
		),
	)

	return c.updateDevice(d, q)
}
