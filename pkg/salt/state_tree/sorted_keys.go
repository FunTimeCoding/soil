package state_tree

import (
	"maps"
	"slices"
)

func sortedKeys(set map[string]bool) []string {
	return slices.Sorted(maps.Keys(set))
}
