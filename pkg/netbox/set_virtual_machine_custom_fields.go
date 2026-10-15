package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetVirtualMachineCustomFields(
	name string,
	fields map[string]any,
) (*virtual_machine.Machine, error) {
	m, e := c.VirtualMachineByName(name)

	if e != nil {
		return nil, e
	}

	q := netbox.NewPatchedWritableVirtualMachineWithConfigContextRequest()
	q.SetName(m.Name)
	q.SetCustomFields(fields)

	return c.updateVirtualMachine(m, q)
}
