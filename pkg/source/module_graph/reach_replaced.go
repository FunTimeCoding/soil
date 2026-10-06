package module_graph

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/build"
	"path/filepath"
	"strings"
)

func reachReplaced(
	c *build.Context,
	nodes map[string]*Node,
	units map[string]*Node,
	replaced map[string]string,
) {
	if len(replaced) == 0 {
		return
	}

	var queue []string

	for _, u := range units {
		queue = append(queue, u.Imports...)
	}

	tried := make(map[string]bool)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if tried[p] {
			continue
		}

		tried[p] = true

		if _, known := nodes[p]; known {
			continue
		}

		for module, directory := range replaced {
			if p != module &&
				!strings.HasPrefix(p, join.Empty(module, constant.Slash)) {
				continue
			}

			relative := filepath.FromSlash(strings.TrimPrefix(p, module))
			n, _ := importDirectory(c, filepath.Join(directory, relative), p)

			if n != nil {
				nodes[p] = n
				queue = append(queue, n.Imports...)
			}
		}
	}
}
