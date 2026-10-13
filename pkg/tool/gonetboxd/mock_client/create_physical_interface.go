package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (c *Client) CreatePhysicalInterface(
	a net.HardwareAddr,
	_ string,
	i *network.Interface,
) (*physical_address.Address, error) {
	c.record("CreatePhysicalInterface", i.Name, a.String())
	v := netbox.NewMACAddressWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetDisplay(a.String())
	v.SetMacAddress(a.String())
	v.SetAssignedObjectType(constant.InterfaceAddress)
	v.SetAssignedObjectId(int64(i.Identifier))

	return physical_address.New(v), nil
}
