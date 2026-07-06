package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/inventory"
	"sync"
)

type Service struct {
	inventory *inventory.Inventory
	sessions  sync.Map
}
