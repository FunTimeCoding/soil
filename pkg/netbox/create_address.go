package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateAddress(
	interfaceIdentifier int32,
	address string,
) (*internet_address.Address, error) {
	q := netbox.NewWritableIPAddressRequest(address)
	q.SetAssignedObjectType(constant.InterfaceAddress)
	q.SetAssignedObjectId(int64(interfaceIdentifier))
	result, _, e := c.client.IpamAPI.IpamIpAddressesCreate(
		c.context,
	).WritableIPAddressRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return internet_address.New(result), nil
}
