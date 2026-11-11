package load

import (
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func ReachUnits(
	g *module_graph.Graph,
	packagePaths []string,
) []*module_graph.Node {
	seen := map[string]bool{}
	var result []*module_graph.Node
	add := func(path string) {
		for _, u := range module_graph.Dependents(g, path) {
			if !seen[u.Path] {
				seen[u.Path] = true
				result = append(result, u)
			}
		}
	}

	for _, p := range packagePaths {
		base, recursive := strings.CutSuffix(p, goModConstant.RecursivePattern)

		if !recursive {
			add(base)

			continue
		}

		prefix := join.Empty(base, "/")

		for path := range g.Nodes {
			if path == base || strings.HasPrefix(path, prefix) {
				add(path)
			}
		}

		for path := range g.Units {
			if path == base || strings.HasPrefix(path, prefix) {
				add(path)
			}
		}
	}

	return result
}
