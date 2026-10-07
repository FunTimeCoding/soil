package service

import "github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"

func (s *Service) Instances() []instance.Instance {
	return s.inventory.Instances
}
