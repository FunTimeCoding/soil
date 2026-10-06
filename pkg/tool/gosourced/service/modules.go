package service

import "github.com/funtimecoding/soil/pkg/source/inventory/module"

func (s *Service) Modules() []module.Module {
	return s.inventory.Modules
}
