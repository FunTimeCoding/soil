package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"net"
)

func (c *Client) FindUnusedAddress(sub *net.IPNet) (net.IP, error) {
	addresses, e := c.InternetAddresses()

	if e != nil {
		return nil, e
	}

	return internet_address.FindUnused(
		internet_address.FilterSubnet(addresses, sub),
		sub,
	), nil
}
