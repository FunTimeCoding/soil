package service

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory"

func (s *Service) Instance(name string) (*inventory.Instance, bool) {
	for i := range s.inventory.Instances {
		if s.inventory.Instances[i].Name == name {
			return &s.inventory.Instances[i], true
		}
	}

	return nil, false
}
