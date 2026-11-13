package gocheck

import (
	"maps"
	"slices"
)

func directoryList(directories map[string]bool) []string {
	return slices.Sorted(maps.Keys(directories))
}
