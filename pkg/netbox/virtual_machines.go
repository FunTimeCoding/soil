package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (c *Client) VirtualMachines() ([]*virtual_machine.Machine, error) {
	result, _, e := c.client.VirtualizationAPI.VirtualizationVirtualMachinesList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return virtual_machine.NewSlice(result.Results), nil
}
