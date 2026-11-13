package gocheck

import (
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"maps"
	"slices"
)

func patternUnits(
	g *module_graph.Graph,
	root string,
	patterns []string,
) []*module_graph.Node {
	var result []*module_graph.Node

	for _, path := range slices.Sorted(maps.Keys(g.Units)) {
		u := g.Units[path]
		relative := relativeDirectory(root, u.Directory)

		for _, p := range patterns {
			if matchesPattern(relative, p) {
				result = append(result, u)

				break
			}
		}
	}

	return result
}
