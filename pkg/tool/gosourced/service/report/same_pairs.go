package report

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"maps"
	"slices"
)

func samePairs(
	before map[string]*symbol.Symbol,
	after map[string]*symbol.Symbol,
) []*entry.Entry {
	var result []*entry.Entry

	for _, key := range slices.Sorted(maps.Keys(before)) {
		to, found := after[key]

		if !found {
			continue
		}

		result = append(result, entry.New("", before[key], to))
		delete(before, key)
		delete(after, key)
	}

	return result
}
