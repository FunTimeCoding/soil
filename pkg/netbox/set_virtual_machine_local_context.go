package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetVirtualMachineLocalContext(
	name string,
	payload map[string]any,
) (*virtual_machine.Machine, error) {
	m, e := c.VirtualMachine(name)

	if e != nil {
		return nil, e
	}

	q := netbox.NewPatchedWritableVirtualMachineWithConfigContextRequest()
	q.SetName(m.Name)
	q.SetLocalContextData(localContext(payload))

	return c.updateVirtualMachine(m, q)
}
