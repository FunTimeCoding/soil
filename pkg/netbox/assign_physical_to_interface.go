package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) AssignPhysicalToInterface(
	p *physical_address.Address,
	i *network.Interface,
) (*physical_address.Address, error) {
	q := netbox.NewMACAddressRequest(p.Name)

	if i == nil {
		q.SetAssignedObjectTypeNil()
		q.SetAssignedObjectIdNil()
	} else {
		q.SetAssignedObjectType(constant.InterfaceAddress)
		q.SetAssignedObjectId(int64(i.Identifier))
	}

	result, _, e := c.client.DcimAPI.DcimMacAddressesUpdate(
		c.context,
		p.Identifier,
	).MACAddressRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return physical_address.New(result), nil
}
