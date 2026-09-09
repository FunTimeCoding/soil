package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

func (c *Client) CreateDevice(
	_ string,
	_ *device_role.Role,
	_ []string,
	_ *device_type.Type,
	_ *site.Site,
	_ *tenant.Tenant,
) (*device.Device, error) {
	return nil, nil
}
