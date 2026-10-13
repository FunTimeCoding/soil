package service

import "github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"

func (s *Service) Inventory() *instance.Inventory {
	return s.inventory
}
