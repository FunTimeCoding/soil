package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) RemoveVirtualTag(
	name string,
	tag string,
) (*virtual_machine.Machine, error) {
	vm, e := c.VirtualMachineByName(name)

	if e != nil {
		return nil, e
	}

	vm.RemoveTag(tag)
	q := netbox.NewPatchedWritableVirtualMachineWithConfigContextRequest()
	q.SetName(vm.Name)
	tags, f := c.tagsNestedRequest(vm.Tags)

	if f != nil {
		return nil, f
	}

	q.SetTags(tags)
	result, _, g := c.client.VirtualizationAPI.VirtualizationVirtualMachinesPartialUpdate(
		c.context,
		vm.Identifier,
	).PatchedWritableVirtualMachineWithConfigContextRequest(*q).Execute()

	if g != nil {
		return nil, g
	}

	return virtual_machine.New(result), nil
}
