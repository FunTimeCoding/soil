package warm_tree

import (
	"maps"
	"slices"
)

func (t *Tree) Files() []string {
	var result []string

	for _, p := range slices.Sorted(maps.Keys(t.entries)) {
		if !t.entries[p].Directory {
			result = append(result, p)
		}
	}

	return result
}
