package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address_range"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateInternetAddressRange(
	start string,
	end string,
	status string,
	description string,
) (*internet_address_range.Range, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*internet_address_range.Range, error) {
			return s.client.CreateInternetAddressRange(
				start,
				end,
				status,
				description,
			)
		},
	)
}
