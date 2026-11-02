package report

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"maps"
	"slices"
)

func packagePairs(
	before map[string]*symbol.Symbol,
	after map[string]*symbol.Symbol,
	pairs []*entry.Entry,
) []*entry.Entry {
	destinations := packageDestinations(pairs)
	added := map[string][]*symbol.Symbol{}

	for _, s := range after {
		key := join.Space(s.Path, s.Member())
		added[key] = append(added[key], s)
	}

	var result []*entry.Entry

	for _, key := range slices.Sorted(maps.Keys(before)) {
		from := before[key]
		targets := destinations[from.Path]

		if len(targets) != 1 {
			continue
		}

		candidates := added[join.Space(targets[0], from.Member())]

		if len(candidates) != 1 {
			continue
		}

		to := candidates[0]
		result = append(result, entry.New("", from, to))
		delete(before, key)
		delete(after, to.Key())
		delete(added, join.Space(to.Path, to.Member()))
	}

	return result
}
