package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) AddVirtualTag(
	name string,
	tag string,
) (*virtual_machine.Machine, error) {
	_, e := c.TagByName(tag)

	if e != nil {
		return nil, e
	}

	vm, f := c.VirtualMachineByName(name)

	if f != nil {
		return nil, f
	}

	vm.AddTag(tag)
	q := netbox.NewPatchedWritableVirtualMachineWithConfigContextRequest()
	q.SetName(vm.Name)
	tags, g := c.tagsNestedRequest(vm.Tags)

	if g != nil {
		return nil, g
	}

	q.SetTags(tags)
	result, _, h := c.client.VirtualizationAPI.VirtualizationVirtualMachinesPartialUpdate(
		c.context,
		vm.Identifier,
	).PatchedWritableVirtualMachineWithConfigContextRequest(*q).Execute()

	if h != nil {
		return nil, h
	}

	return virtual_machine.New(result), nil
}
