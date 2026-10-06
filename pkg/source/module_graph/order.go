package module_graph

import "sort"

func order(nodes map[string]*Node) []string {
	paths := make([]string, 0, len(nodes))

	for p := range nodes {
		paths = append(paths, p)
	}

	sort.Strings(paths)
	visited := make(map[string]bool, len(nodes))
	result := make([]string, 0, len(nodes))
	var visit func(string)
	visit = func(p string) {
		if visited[p] {
			return
		}

		visited[p] = true

		for _, i := range nodes[p].Imports {
			if _, okay := nodes[i]; okay {
				visit(i)
			}
		}

		result = append(result, p)
	}

	for _, p := range paths {
		visit(p)
	}

	return result
}
