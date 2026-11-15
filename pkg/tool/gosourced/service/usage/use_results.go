package usage

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
	"maps"
	"slices"
)

func useResults(uses map[string]map[string]bool) []*result.UsedType {
	labels := slices.Sorted(maps.Keys(uses))
	entries := make([]*result.UsedType, 0, len(labels))

	for _, l := range labels {
		entries = append(entries, result.NewUsedType(l, sortedKeys(uses[l])))
	}

	return entries
}
