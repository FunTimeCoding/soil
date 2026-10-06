package service

import (
	"github.com/funtimecoding/soil/pkg/source/index/cache"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"sync"
)

type Service struct {
	inventory    *inventory.Inventory
	sessions     sync.Map
	commits      sync.Mutex
	beforeCommit func()
	workspaces   *cache.Cache
	full         bool
}
