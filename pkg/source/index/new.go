package index

import (
	"github.com/funtimecoding/soil/pkg/source/index/store"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
)

func New(
	directory string,
	root string,
	kinds ...*Kind,
) *Workspace {
	return &Workspace{
		root:  root,
		graph: module_graph.Build(root),
		store: store.New(directory),
		kinds: kinds,
		sums:  make(map[string]*fileSum),
		memo:  make(map[string]*memoEntry),
	}
}
