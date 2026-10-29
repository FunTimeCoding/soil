package service

import "github.com/funtimecoding/soil/pkg/netbox/cable"

func (s *Service) Cables() ([]*cable.Cable, error) {
	return s.client.Cables()
}
