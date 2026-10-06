package module_graph

import "sort"

func Dependents(
	g *Graph,
	path string,
) []*Node {
	reach := map[string]bool{path: true}

	for _, p := range g.Order {
		if reaches(g.Nodes[p].Imports, reach) {
			reach[p] = true
		}
	}

	var result []*Node

	for _, u := range g.Units {
		if reach[u.Path] || reaches(u.Imports, reach) {
			result = append(result, u)
		}
	}

	sort.Slice(
		result,
		func(
			i int,
			j int,
		) bool {
			return result[i].Path < result[j].Path
		},
	)

	return result
}
