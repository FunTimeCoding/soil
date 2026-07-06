package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) UpdatePhysicalAddress(
	a *physical_address.Address,
	i *network.Interface,
) (*physical_address.Address, error) {
	q := netbox.NewMACAddressRequest(a.Name)
	q.SetAssignedObjectType(constant.InterfaceAddress)
	q.SetAssignedObjectId(int64(i.Identifier))
	result, _, e := c.client.DcimAPI.DcimMacAddressesUpdate(
		c.context,
		a.Identifier,
	).MACAddressRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return physical_address.New(result), nil
}
