package workspace

import "github.com/funtimecoding/soil/pkg/source/module_graph"

func (w *Workspace) Commit() ([]string, error) {
	if len(w.layer.Changes()) == 0 {
		return nil, nil
	}

	graphs := make([]*module_graph.Graph, 0, len(w.roots))

	for _, root := range w.roots {
		graphs = append(graphs, module_graph.Build(root))
	}

	return w.layer.Commit(graphs)
}
