package census

import (
	"maps"
	"slices"
)

func sortedKeys(m map[string]bool) []string {
	return slices.Sorted(maps.Keys(m))
}
