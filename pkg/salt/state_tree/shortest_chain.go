package state_tree

import "slices"

func (t *Tree) shortestChain(
	from string,
	to string,
) []string {
	parent := map[string]string{from: ""}
	queue := []string{from}

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		if s == to {
			var result []string

			for c := s; c != ""; c = parent[c] {
				result = append(result, c)
			}

			slices.Reverse(result)

			return result
		}

		for _, n := range t.includes[s] {
			if _, seen := parent[n]; seen {
				continue
			}

			parent[n] = s
			queue = append(queue, n)
		}
	}

	return nil
}
