package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/physical_address"

func (c *Client) AssignPhysicalToVirtualInterface(
	_ *physical_address.Address,
	_ int32,
) (*physical_address.Address, error) {
	return nil, c.Failure
}
