package reach

import "github.com/funtimecoding/soil/pkg/source/module_graph"

func NewReach(
	packages []string,
	documents []string,
	graph *module_graph.Graph,
) *Reach {
	return &Reach{Packages: packages, Documents: documents, Graph: graph}
}
