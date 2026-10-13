package mock_service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"

func (s *Service) Inventory() *instance.Inventory {
	return s.inventory
}
