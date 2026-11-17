package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (c *Client) SetVirtualMachineLocalContext(
	name string,
	payload map[string]any,
) (*virtual_machine.Machine, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	m, e := c.VirtualMachine(name)

	if e != nil {
		return nil, e
	}

	m.LocalContext = nil

	if len(payload) > 0 {
		m.LocalContext = payload
	}

	c.record("SetVirtualMachineLocalContext", name)

	return m, nil
}
