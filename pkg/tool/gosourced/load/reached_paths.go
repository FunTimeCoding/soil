package load

import "github.com/funtimecoding/soil/pkg/source/module_graph"

func ReachedPaths(
	g *module_graph.Graph,
	packagePaths []string,
) map[string]bool {
	result := map[string]bool{}

	for _, u := range module_graph.ReachUnits(g, packagePaths) {
		result[u.Path] = true
	}

	return result
}
