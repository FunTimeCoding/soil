package service

import "github.com/funtimecoding/soil/pkg/source/module_graph"

func reachedPaths(
	g *module_graph.Graph,
	packagePaths []string,
) map[string]bool {
	result := map[string]bool{}

	for _, u := range reachUnits(g, packagePaths) {
		result[u.Path] = true
	}

	return result
}
