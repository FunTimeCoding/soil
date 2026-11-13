package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address_range"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) InternetAddressRanges() ([]*internet_address_range.Range, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*internet_address_range.Range, error) {
			return s.client.InternetAddressRanges()
		},
	)
}
