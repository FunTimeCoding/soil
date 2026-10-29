package service

import "github.com/funtimecoding/soil/pkg/netbox/location"

func (s *Service) Locations() ([]*location.Location, error) {
	return s.client.Locations()
}
