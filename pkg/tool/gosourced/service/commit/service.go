package commit

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn"
	"sync"
)

type Service struct {
	inventory    *inventory.Inventory
	workspace    *workspace.Service
	commits      sync.Mutex
	beforeCommit func()
	memory       face.Client
	guard        *peer.Guard
	reporter     soilFace.Reporter
	turns        *turn.Scheduler
}
