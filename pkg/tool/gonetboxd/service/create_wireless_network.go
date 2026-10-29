package service

import "github.com/funtimecoding/soil/pkg/netbox/wireless_network"

func (s *Service) CreateWirelessNetwork(ssid string) (*wireless_network.Network, error) {
	return s.client.CreateWirelessNetwork(ssid)
}
