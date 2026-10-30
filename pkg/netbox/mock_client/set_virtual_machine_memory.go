package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (c *Client) SetVirtualMachineMemory(
	name string,
	memory int32,
) (*virtual_machine.Machine, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	m, e := c.VirtualMachineByName(name)

	if e != nil {
		return nil, e
	}

	m.Raw.SetMemory(memory)

	return m, nil
}
