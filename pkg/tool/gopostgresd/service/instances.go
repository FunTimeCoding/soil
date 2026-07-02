package service

import "github.com/funtimecoding/go-library/pkg/tool/gopostgresd/inventory"

func (s *Service) Instances() []inventory.Instance {
	return s.inventory.Instances
}
