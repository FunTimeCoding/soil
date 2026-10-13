package service

import "github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"

func (s *Service) Inventory() *instance.Inventory {
	return s.inventory
}
