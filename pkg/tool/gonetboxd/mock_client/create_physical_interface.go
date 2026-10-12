package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"net"
)

func (c *Client) CreatePhysicalInterface(
	a net.HardwareAddr,
	_ string,
	i *network.Interface,
) (*physical_address.Address, error) {
	c.record("CreatePhysicalInterface", i.Name, a.String())

	return nil, nil
}
