package service

import "github.com/funtimecoding/soil/pkg/netbox/manufacturer"

func (s *Service) CreateManufacturer(name string) (*manufacturer.Manufacturer, error) {
	return s.client.CreateManufacturer(name)
}
