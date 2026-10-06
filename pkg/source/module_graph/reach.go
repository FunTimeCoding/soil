package module_graph

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/build"
	"path/filepath"
	"strings"
)

func reach(
	c *build.Context,
	module string,
	nodes map[string]*Node,
	units map[string]*Node,
	replaced map[string]string,
) {
	var queue []string

	for _, u := range units {
		queue = append(queue, u.Imports...)
	}

	visited := make(map[string]bool)
	reached := make(map[string]bool)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if visited[p] {
			continue
		}

		visited[p] = true

		if p == module || strings.HasPrefix(p, join.Empty(module, "/")) {
			continue
		}

		n, known := nodes[p]

		if !known {
			for m, directory := range replaced {
				if p != m && !strings.HasPrefix(p, join.Empty(m, "/")) {
					continue
				}

				relative := filepath.FromSlash(strings.TrimPrefix(p, m))
				n, _ = importDirectory(c, filepath.Join(directory, relative), p)
			}
		}

		if n == nil {
			continue
		}

		nodes[p] = n
		reached[p] = true
		queue = append(queue, n.Imports...)
	}

	for p := range nodes {
		if p != module && !strings.HasPrefix(p, join.Empty(module, "/")) &&
			!reached[p] {
			delete(nodes, p)
		}
	}
}
