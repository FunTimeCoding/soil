package mock_service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"

func (s *Service) index(name string) int {
	return instance.Index(s.inventory, name)
}
