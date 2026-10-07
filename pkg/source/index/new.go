package index

import (
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"github.com/funtimecoding/soil/pkg/source/index/store"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/source/types/file_sum"
	"github.com/funtimecoding/soil/pkg/source/types/memo_entry"
)

func New(
	directory string,
	root string,
	kinds ...*kind.Kind,
) *Workspace {
	return &Workspace{
		root:  root,
		graph: module_graph.Build(root),
		store: store.New(directory),
		kinds: kinds,
		sums:  make(map[string]*file_sum.Sum),
		memo:  make(map[string]*memo_entry.Entry),
	}
}
