package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (c *Client) AddVirtualTag(
	name string,
	tag string,
) (*virtual_machine.Machine, error) {
	return c.changeMachine("AddVirtualTag", name, tag)
}
