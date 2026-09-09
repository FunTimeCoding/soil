package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (c *Client) RenameVirtualMachine(
	_ string,
	_ string,
) (*virtual_machine.Machine, error) {
	return nil, nil
}
