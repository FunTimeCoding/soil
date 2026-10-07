package index

import (
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"github.com/funtimecoding/soil/pkg/source/index/store"
	"github.com/funtimecoding/soil/pkg/source/index/xref"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/source/types/file_sum"
	"github.com/funtimecoding/soil/pkg/source/types/memo_entry"
	"sync"
)

type Workspace struct {
	root         string
	graph        *module_graph.Graph
	store        *store.Store
	kinds        []*kind.Kind
	lock         sync.Mutex
	sums         map[string]*file_sum.Sum
	memo         map[string]*memo_entry.Entry
	fingerprints map[string]string
	facts        map[string]map[string]any
	references   *xref.Index
}
