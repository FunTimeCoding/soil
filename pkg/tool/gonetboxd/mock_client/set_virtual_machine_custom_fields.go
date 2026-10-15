package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"maps"
)

func (c *Client) SetVirtualMachineCustomFields(
	name string,
	fields map[string]any,
) (*virtual_machine.Machine, error) {
	m, e := c.VirtualMachineByName(name)

	if e != nil {
		return nil, e
	}

	if m.CustomFields == nil {
		m.CustomFields = map[string]any{}
	}

	maps.Copy(m.CustomFields, fields)
	c.record("SetVirtualMachineCustomFields", name)

	return m, nil
}
