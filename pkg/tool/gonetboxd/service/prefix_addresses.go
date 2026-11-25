package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/peer"
	"net"
)

func (s *Service) PrefixAddresses(prefix string) ([]*internet_address.Address, error) {
	if _, _, e := net.ParseCIDR(prefix); e != nil {
		return nil, validation.New("prefix %s: not in CIDR notation", prefix)
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*internet_address.Address, error) {
			return s.client.PrefixAddresses(prefix)
		},
	)
}
