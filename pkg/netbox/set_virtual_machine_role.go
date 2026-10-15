package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetVirtualMachineRole(
	name string,
	roleName string,
) (*virtual_machine.Machine, error) {
	m, e := c.VirtualMachineByName(name)

	if e != nil {
		return nil, e
	}

	r, f := c.DeviceRoleByName(roleName)

	if f != nil {
		return nil, f
	}

	q := netbox.NewPatchedWritableVirtualMachineWithConfigContextRequest()
	q.SetName(m.Name)
	q.SetRole(
		netbox.PatchedWritableVirtualMachineWithConfigContextRequestRole{
			Int32: &r.Identifier,
		},
	)

	return c.updateVirtualMachine(m, q)
}
