package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (c *Client) MustVirtualMachines() []*virtual_machine.Machine {
	result, e := c.VirtualMachines()
	errors.PanicOnError(e)

	return result
}
