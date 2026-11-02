package report

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"maps"
	"slices"
)

func namePairs(
	before map[string]*symbol.Symbol,
	after map[string]*symbol.Symbol,
) []*entry.Entry {
	removed := byMember(before)
	added := byMember(after)
	var result []*entry.Entry

	for _, member := range slices.Sorted(maps.Keys(removed)) {
		from := removed[member]
		to := added[member]

		if len(from) != 1 || len(to) != 1 {
			continue
		}

		result = append(result, entry.New("", from[0], to[0]))
		delete(before, from[0].Key())
		delete(after, to[0].Key())
	}

	return result
}
