package service

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"

func (s *Service) Instances() []instance.Instance {
	return s.inventory.Instances
}
