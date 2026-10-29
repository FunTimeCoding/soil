package service

import "github.com/funtimecoding/soil/pkg/netbox/tunnel_group"

func (s *Service) CreateTunnelGroup(name string) (*tunnel_group.Group, error) {
	return s.client.CreateTunnelGroup(name)
}
