package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (c *Client) MustVirtualMachineByName(n string) *virtual_machine.Machine {
	result, e := c.VirtualMachineByName(n)
	errors.PanicOnError(e)

	return result
}
