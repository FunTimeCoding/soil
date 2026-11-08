package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
	"maps"
	"slices"
)

func useResults(uses map[string]map[string]bool) []*result.CensusUse {
	labels := slices.Sorted(maps.Keys(uses))
	entries := make([]*result.CensusUse, 0, len(labels))

	for _, l := range labels {
		entries = append(entries, result.NewCensusUse(l, sortedKeys(uses[l])))
	}

	return entries
}
