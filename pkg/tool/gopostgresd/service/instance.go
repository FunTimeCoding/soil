package service

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"

func (s *Service) Instance(name string) (*instance.Instance, bool) {
	for i := range s.inventory.Instances {
		if s.inventory.Instances[i].Name == name {
			return &s.inventory.Instances[i], true
		}
	}

	return nil, false
}
