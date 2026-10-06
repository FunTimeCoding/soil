package service

import "github.com/funtimecoding/soil/pkg/source/inventory/module"

func (s *Service) Module(name string) (*module.Module, bool) {
	for _, m := range s.inventory.Modules {
		if m.Name == name {
			return &m, true
		}
	}

	return nil, false
}
