package diff

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module"
)

func New(
	i *inventory.Inventory,
	m *module.Service,
) *Service {
	return &Service{inventory: i, module: m}
}
