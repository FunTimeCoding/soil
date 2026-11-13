package workspace

import (
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"slices"
)

func censusPatterns(
	g *module_graph.Graph,
	packagePaths []string,
	full bool,
) ([]string, bool) {
	whole := []string{goModConstant.AllPackages}
	var patterns []string

	for _, packagePath := range packagePaths {
		_, own := g.Units[packagePath]
		_, reached := g.Nodes[packagePath]

		if own || reached {
			patterns = append(patterns, packagePath)
		}

		if reached && !own {
			whole = append(whole, packagePath)
		}

		for _, u := range module_graph.Dependents(g, packagePath) {
			patterns = append(patterns, u.Directory)
		}
	}

	if full {
		return whole, true
	}

	patterns = slices.Compact(slices.Sorted(slices.Values(patterns)))

	if 2*len(patterns) > len(g.Units) {
		return whole, true
	}

	return patterns, false
}
