package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"maps"
	"net"
	"slices"
)

func (c *Client) PrefixAddresses(prefix string) ([]*internet_address.Address, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	_, network, e := net.ParseCIDR(prefix)

	if e != nil {
		return nil, validation.New("prefix %s: not in CIDR notation", prefix)
	}

	result := []*internet_address.Address{}

	for _, host := range slices.Sorted(maps.Keys(c.addresses)) {
		result = append(
			result,
			internet_address.FilterSubnet(c.addresses[host], network)...,
		)
	}

	return result, nil
}
