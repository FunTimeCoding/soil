package report

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"maps"
	"slices"
)

func receiverPairs(
	before map[string]*symbol.Symbol,
	after map[string]*symbol.Symbol,
	pairs []*entry.Entry,
) []*entry.Entry {
	destinations := receiverDestinations(pairs)
	added := map[string]*symbol.Symbol{}

	for _, s := range after {
		if s.Kind == constant.SymbolType {
			added[s.Reference()] = s
		}
	}

	var result []*entry.Entry

	for _, key := range slices.Sorted(maps.Keys(before)) {
		from := before[key]

		if from.Kind != constant.SymbolType {
			continue
		}

		targets := destinations[from.Reference()]

		if len(targets) != 1 {
			continue
		}

		to, found := added[targets[0]]

		if !found {
			continue
		}

		result = append(result, entry.New("", from, to))
		delete(before, key)
		delete(after, to.Key())
		delete(added, to.Reference())
	}

	return result
}
