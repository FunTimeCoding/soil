package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetVirtualMachineTenant(
	name string,
	tenantName string,
) (*virtual_machine.Machine, error) {
	m, e := c.VirtualMachineByName(name)

	if e != nil {
		return nil, e
	}

	n, f := c.TenantByName(tenantName)

	if f != nil {
		return nil, f
	}

	q := netbox.NewPatchedWritableVirtualMachineWithConfigContextRequest()
	q.SetName(m.Name)
	q.SetTenant(
		netbox.BriefTenantRequestAsASNRangeRequestTenant(
			netbox.NewBriefTenantRequest(n.Name, n.Raw.Slug),
		),
	)

	return c.updateVirtualMachine(m, q)
}
