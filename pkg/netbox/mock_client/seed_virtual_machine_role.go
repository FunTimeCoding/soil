package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedVirtualMachineRole(name string) {
	v := netbox.NewDeviceRoleWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetName(name)
	v.SetVmRole(true)
	c.roles = append(c.roles, device_role.New(v))
}
