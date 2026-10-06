package index

import "github.com/funtimecoding/soil/pkg/source/module_graph"

func ready(
	g *module_graph.Graph,
	path string,
	fingerprints map[string]string,
) bool {
	for _, i := range g.Nodes[path].Imports {
		if _, workspace := g.Nodes[i]; !workspace {
			continue
		}

		if _, known := fingerprints[i]; !known {
			return false
		}
	}

	return true
}
