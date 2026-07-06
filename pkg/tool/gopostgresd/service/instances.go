package service

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory"

func (s *Service) Instances() []inventory.Instance {
	return s.inventory.Instances
}
