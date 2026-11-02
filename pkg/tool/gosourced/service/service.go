package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/source/index/cache"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	memoryFace "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/turn"
	"sync"
)

type Service struct {
	inventory    *inventory.Inventory
	commits      sync.Mutex
	beforeCommit func()
	workspaces   *cache.Cache
	full         bool
	memory       memoryFace.Client
	guard        *peer.Guard
	reporter     face.Reporter
	closeOnce    sync.Once
	turns        *turn.Scheduler
}
