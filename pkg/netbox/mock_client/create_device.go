package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

func (c *Client) CreateDevice(
	name string,
	_ *device_role.Role,
	tags []string,
	_ *device_type.Type,
	_ *site.Site,
	n *tenant.Tenant,
) (*device.Device, error) {
	result := newDevice(c.nextIdentifier(), name, tags, n)
	c.devices = append(c.devices, result)

	return result, nil
}
