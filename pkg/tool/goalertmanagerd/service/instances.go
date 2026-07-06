package service

import "github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/inventory"

func (s *Service) Instances() []inventory.Instance {
	return s.inventory.Instances
}
