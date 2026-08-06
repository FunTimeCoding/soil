package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (c *Client) VirtualMachineInterfaces(
	m *virtual_machine.Machine,
) ([]*network.Interface, error) {
	result, _, e := c.client.VirtualizationAPI.VirtualizationInterfacesList(
		c.context,
	).VirtualMachineId([]int32{m.Identifier}).Execute()

	if e != nil {
		return nil, e
	}

	return network.NewVirtualSlice(result.Results), nil
}
