package reach

import "github.com/funtimecoding/soil/pkg/source/module_graph"

type Reach struct {
	Packages  []string
	Documents []string
	Graph     *module_graph.Graph
}
