package index

import (
	"github.com/funtimecoding/soil/pkg/source/index/store"
	"github.com/funtimecoding/soil/pkg/source/index/xref"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"sync"
)

type Workspace struct {
	root         string
	graph        *module_graph.Graph
	store        *store.Store
	kinds        []*Kind
	lock         sync.Mutex
	sums         map[string]*fileSum
	memo         map[string]*memoEntry
	fingerprints map[string]string
	facts        map[string]map[string]any
	references   *xref.Index
}
