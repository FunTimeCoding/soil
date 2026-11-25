package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"net"
)

func (c *Client) ListPrefixAddresses(
	_ context.Context,
	prefix string,
) ([]*client.Address, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	_, network, e := net.ParseCIDR(prefix)

	if e != nil {
		return nil, validation.New("prefix %s: not in CIDR notation", prefix)
	}

	result := []*client.Address{}

	for _, a := range c.AddressList {
		if network.Contains(net.ParseIP(a.Address)) {
			result = append(result, a)
		}
	}

	return result, nil
}
