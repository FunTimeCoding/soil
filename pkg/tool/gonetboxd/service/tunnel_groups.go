package service

import "github.com/funtimecoding/soil/pkg/netbox/tunnel_group"

func (s *Service) TunnelGroups() ([]*tunnel_group.Group, error) {
	return guarded(
		s.guard,
		func() ([]*tunnel_group.Group, error) {
			return s.client.TunnelGroups()
		},
	)
}
