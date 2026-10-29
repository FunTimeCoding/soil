package service

import "github.com/funtimecoding/soil/pkg/netbox/wireless_network"

func (s *Service) WirelessNetworks() ([]*wireless_network.Network, error) {
	return s.client.WirelessNetworks()
}
