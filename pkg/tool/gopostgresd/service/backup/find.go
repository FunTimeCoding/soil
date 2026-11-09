package backup

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
)

func (s *Service) find(name string) (*instance.Instance, error) {
	result, found := s.inventory.Find(name)

	if !found {
		return nil, not_found.New(inventoryConstant.InstanceKind, name)
	}

	return result, nil
}
