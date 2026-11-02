package commit

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	memoryFace "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn"
	"sync"
)

type Service struct {
	inventory    *inventory.Inventory
	workspace    *workspace.Service
	commits      sync.Mutex
	beforeCommit func()
	memory       memoryFace.Client
	guard        *peer.Guard
	reporter     face.Reporter
	turns        *turn.Scheduler
}
