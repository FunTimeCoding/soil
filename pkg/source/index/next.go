package index

import (
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/system"
	"maps"
)

func Next(
	previous *Workspace,
	changed []string,
) *Workspace {
	previous.lock.Lock()
	sums := maps.Clone(previous.sums)
	memo := maps.Clone(previous.memo)
	previous.lock.Unlock()

	for _, p := range changed {
		if !system.FileExists(p) {
			delete(sums, p)
		}
	}

	return &Workspace{
		root:  previous.root,
		graph: module_graph.Update(previous.graph, previous.root, changed),
		store: previous.store,
		kinds: previous.kinds,
		sums:  sums,
		memo:  memo,
		build: previous.build,
	}
}
