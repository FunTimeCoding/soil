package mock_service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"

func (s *Service) Instances() []inventory.Instance {
	return s.inventory.Instances
}
