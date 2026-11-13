package option

import (
	"github.com/funtimecoding/soil/pkg/lint/document_cache"
	"github.com/funtimecoding/soil/pkg/lint/pointer_cache"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"sync"
)

type Check struct {
	Root         string
	Fix          bool
	Index        string
	Inventory    *inventory.Inventory
	Transcript   string
	Fresh        bool
	Trees        *warm_forest.Forest
	Documents    *document_cache.Cache
	Resolved     *pointer_cache.Cache
	Workspace    func(string) *index.Workspace
	CommitLock   sync.Locker
	BeforeCommit func()
}
