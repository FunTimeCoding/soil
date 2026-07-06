package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (c *Client) MustRemoveVirtualTag(
	name string,
	tag string,
) *virtual_machine.Machine {
	result, e := c.RemoveVirtualTag(name, tag)
	errors.PanicOnError(e)

	return result
}
