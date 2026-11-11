package relocator

import (
	"maps"
	"slices"
)

func cyclePath(
	edges map[string]map[string]bool,
	start string,
) []string {
	parent := map[string]string{start: ""}
	queue := []string{start}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		for _, m := range slices.Sorted(maps.Keys(edges[n])) {
			if m == start {
				var result []string

				for c := n; c != ""; c = parent[c] {
					result = append([]string{c}, result...)
				}

				return append(result, start)
			}

			if _, visited := parent[m]; visited {
				continue
			}

			parent[m] = n
			queue = append(queue, m)
		}
	}

	return nil
}
