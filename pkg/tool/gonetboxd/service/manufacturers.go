package service

import "github.com/funtimecoding/soil/pkg/netbox/manufacturer"

func (s *Service) Manufacturers() ([]*manufacturer.Manufacturer, error) {
	return guarded(
		s.guard,
		func() ([]*manufacturer.Manufacturer, error) {
			return s.client.Manufacturers()
		},
	)
}
