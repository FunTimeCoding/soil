package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gobandd/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory"
	"sync"
)

type Service struct {
	inventory *inventory.Inventory
	engines   map[string]face.Engine
	sessions  sync.Map
	mutex     sync.Mutex
}
