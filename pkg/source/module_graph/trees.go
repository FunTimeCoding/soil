package module_graph

import (
	"maps"
	"slices"
)

func Trees(root string) []string {
	_, _, replaced := readModule(root)

	return append([]string{root}, slices.Sorted(maps.Values(replaced))...)
}
