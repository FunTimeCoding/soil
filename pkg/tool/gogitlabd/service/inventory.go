package service

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"

func (s *Service) Inventory() *instance.Inventory {
	return s.inventory
}
