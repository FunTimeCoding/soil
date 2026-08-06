package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) updateVirtualMachine(
	m *virtual_machine.Machine,
	q *netbox.PatchedWritableVirtualMachineWithConfigContextRequest,
) (*virtual_machine.Machine, error) {
	result, _, e := c.client.VirtualizationAPI.VirtualizationVirtualMachinesPartialUpdate(
		c.context,
		m.Identifier,
	).PatchedWritableVirtualMachineWithConfigContextRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return virtual_machine.New(result), nil
}
