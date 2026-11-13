package index

import (
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
)

func New(
	directory string,
	root string,
	kinds ...*kind.Kind,
) *Workspace {
	return NewWithGraph(directory, root, module_graph.Build(root), kinds...)
}
