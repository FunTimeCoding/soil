package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

// CreatePhysicalInterface Create MAC address and assign to existing interface
func (c *Client) CreatePhysicalInterface(
	a net.HardwareAddr,
	description string,
	i *network.Interface,
) (*physical_address.Address, error) {
	q := netbox.NewMACAddressRequest(a.String())

	if description != "" {
		q.SetDescription(description)
	}

	q.SetAssignedObjectType(constant.InterfaceAddress)
	q.SetAssignedObjectId(int64(i.Identifier))
	result, _, e := c.client.DcimAPI.DcimMacAddressesCreate(
		c.context,
	).MACAddressRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return physical_address.New(result), nil
}
