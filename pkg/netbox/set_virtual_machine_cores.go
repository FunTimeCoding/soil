package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetVirtualMachineCores(
	name string,
	cores float64,
) (*virtual_machine.Machine, error) {
	m, e := c.VirtualMachineByName(name)

	if e != nil {
		return nil, e
	}

	q := netbox.NewPatchedWritableVirtualMachineWithConfigContextRequest()
	q.SetName(m.Name)
	q.SetVcpus(cores)

	return c.updateVirtualMachine(m, q)
}
