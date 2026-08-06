package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (c *Client) VirtualMachineDisks(
	m *virtual_machine.Machine,
) ([]*virtual_disk.Disk, error) {
	result, _, e := c.client.VirtualizationAPI.VirtualizationVirtualDisksList(
		c.context,
	).VirtualMachineId([]int32{m.Identifier}).Execute()

	if e != nil {
		return nil, e
	}

	return virtual_disk.NewSlice(result.Results), nil
}
