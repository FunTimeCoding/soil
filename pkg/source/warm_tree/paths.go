package warm_tree

import (
	"maps"
	"slices"
)

func (t *Tree) Paths() []string {
	return slices.Sorted(maps.Keys(t.entries))
}
