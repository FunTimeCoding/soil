package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedHostAddress(
	host string,
	address string,
	objectType string,
	objectIdentifier int64,
) *internet_address.Address {
	v := netbox.NewIPAddressWithDefaults()
	v.SetAddress(address)
	v.SetDisplay(address)
	v.SetAssignedObjectType(objectType)
	v.SetAssignedObjectId(objectIdentifier)
	result := internet_address.New(v)
	c.addresses[host] = append(c.addresses[host], result)

	return result
}
