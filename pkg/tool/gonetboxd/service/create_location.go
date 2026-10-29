package service

import "github.com/funtimecoding/soil/pkg/netbox/location"

func (s *Service) CreateLocation(
	name string,
	siteName string,
) (*location.Location, error) {
	return s.client.CreateLocation(name, siteName)
}
