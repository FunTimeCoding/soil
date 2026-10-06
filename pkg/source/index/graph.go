package index

import "github.com/funtimecoding/soil/pkg/source/module_graph"

func (w *Workspace) Graph() *module_graph.Graph {
	return w.graph
}
