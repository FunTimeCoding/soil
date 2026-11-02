package module

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
)

func New(
	i *inventory.Inventory,
	c *commit.Service,
) *Service {
	return &Service{inventory: i, commit: c}
}
