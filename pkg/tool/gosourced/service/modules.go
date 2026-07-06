package service

import "github.com/funtimecoding/soil/pkg/tool/gosourced/inventory"

func (s *Service) Modules() []inventory.Module {
	return s.inventory.Modules
}
