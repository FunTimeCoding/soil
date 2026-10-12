package service

import "github.com/funtimecoding/soil/pkg/netbox/tunnel"

func (s *Service) CreateTunnel(
	name string,
	encapsulation string,
	group string,
) (*tunnel.Tunnel, error) {
	g, e := s.client.TunnelGroupByName(group)

	if e != nil {
		return nil, e
	}

	return s.client.CreateTunnel(name, encapsulation, g)
}
