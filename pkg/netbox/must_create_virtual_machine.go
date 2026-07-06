package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (c *Client) MustCreateVirtualMachine(
	name string,
	cl *cluster.Cluster,
) *virtual_machine.Machine {
	result, e := c.CreateVirtualMachine(name, cl)
	errors.PanicOnError(e)

	return result
}
