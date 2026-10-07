package service

import "github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"

func (s *Service) Instances() []instance.Instance {
	return s.inventory.Instances
}
