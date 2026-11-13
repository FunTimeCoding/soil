package layer

import (
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"path/filepath"
)

func packagePaths(
	graphs []*module_graph.Graph,
	touched []string,
) []string {
	directories := map[string]bool{}

	for _, path := range touched {
		directories[filepath.Dir(path)] = true
	}

	var result []string

	for _, g := range graphs {
		for path, n := range g.Nodes {
			if directories[n.Directory] {
				result = append(result, path)
			}
		}
	}

	return result
}
