package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (c *Client) CreatePhysical(
	a net.HardwareAddr,
	description string,
) (*physical_address.Address, error) {
	q := netbox.NewMACAddressRequest(a.String())

	if description != "" {
		q.SetDescription(description)
	}

	result, _, e := c.client.DcimAPI.DcimMacAddressesCreate(
		c.context,
	).MACAddressRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return physical_address.New(result), nil
}
