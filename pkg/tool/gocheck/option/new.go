package option

import (
	"github.com/funtimecoding/soil/pkg/lint/document_cache"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"sync"
)

func New() *Check {
	return &Check{
		Trees:      warm_forest.New(),
		Documents:  document_cache.New(),
		CommitLock: new(sync.Mutex),
	}
}
