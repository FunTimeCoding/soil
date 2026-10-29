package service

import "github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"

func (s *Service) TunnelTerminations() ([]*tunnel_termination.Termination, error) {
	return s.client.TunnelTerminations()
}
