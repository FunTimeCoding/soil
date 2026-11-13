package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateWirelessNetwork(ssid string) (*wireless_network.Network, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*wireless_network.Network, error) {
			return s.client.CreateWirelessNetwork(ssid)
		},
	)
}
