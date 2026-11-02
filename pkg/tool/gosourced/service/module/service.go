package module

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
)

type Service struct {
	inventory *inventory.Inventory
	commit    *commit.Service
}
