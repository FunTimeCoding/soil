package module_graph

import "github.com/funtimecoding/soil/pkg/source/warm_forest"

func Build(root string) *Graph {
	return BuildForest(warm_forest.New(), root)
}
