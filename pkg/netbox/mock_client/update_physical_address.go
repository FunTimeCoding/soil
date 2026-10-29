package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
)

func (c *Client) UpdatePhysicalAddress(
	_ *physical_address.Address,
	_ *network.Interface,
) (*physical_address.Address, error) {
	return nil, c.Failure
}
