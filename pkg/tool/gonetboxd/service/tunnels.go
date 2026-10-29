package service

import "github.com/funtimecoding/soil/pkg/netbox/tunnel"

func (s *Service) Tunnels() ([]*tunnel.Tunnel, error) {
	return s.client.Tunnels()
}
