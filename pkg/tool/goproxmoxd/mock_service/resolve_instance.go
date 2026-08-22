package mock_service

import "github.com/funtimecoding/soil/pkg/inventory"

func (s *Service) ResolveInstance(explicit string) (string, error) {
	return inventory.Resolve(explicit, s.inventory.Names())
}
