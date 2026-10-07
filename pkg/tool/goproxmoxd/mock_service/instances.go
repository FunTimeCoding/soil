package mock_service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"

func (s *Service) Instances() []instance.Instance {
	return s.inventory.Instances
}
