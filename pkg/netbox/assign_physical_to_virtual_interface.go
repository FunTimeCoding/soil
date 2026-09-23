package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) AssignPhysicalToVirtualInterface(
	p *physical_address.Address,
	identifier int32,
) (*physical_address.Address, error) {
	q := netbox.NewMACAddressRequest(p.Name)
	q.SetAssignedObjectType(constant.VirtualInterfaceAddress)
	q.SetAssignedObjectId(int64(identifier))
	result, _, e := c.client.DcimAPI.DcimMacAddressesUpdate(
		c.context,
		p.Identifier,
	).MACAddressRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	c.cache.PhysicalAddresses = nil

	return physical_address.New(result), nil
}
