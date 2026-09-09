package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (c *Client) SetVirtualMachineMemory(
	_ string,
	_ int32,
) (*virtual_machine.Machine, error) {
	return nil, nil
}
