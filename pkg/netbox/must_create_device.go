package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

func (c *Client) MustCreateDevice(
	name string,
	o *device_role.Role,
	tags []string,
	y *device_type.Type,
	s *site.Site,
	n *tenant.Tenant,
) *device.Device {
	result, e := c.CreateDevice(name, o, tags, y, s, n)
	errors.PanicOnError(e)

	return result
}
