package index

import "github.com/funtimecoding/soil/pkg/source/module_graph"

func remaining(
	g *module_graph.Graph,
	done map[string]bool,
	failed map[string]bool,
) []string {
	var paths []string

	for _, path := range g.Order {
		if !done[path] && !failed[path] {
			paths = append(paths, path)
		}
	}

	return paths
}
