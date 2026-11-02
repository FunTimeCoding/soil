package report

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"maps"
	"slices"
)

func shapePairs(
	before map[string]*symbol.Symbol,
	after map[string]*symbol.Symbol,
	mapping map[string]string,
) []*entry.Entry {
	shapes := map[string][]*symbol.Symbol{}

	for _, key := range slices.Sorted(maps.Keys(after)) {
		s := after[key]
		shapes[s.Shape] = append(shapes[s.Shape], s)
	}

	var result []*entry.Entry

	for _, key := range slices.Sorted(maps.Keys(before)) {
		from := before[key]
		shape := remap(from.Shape, mapping)

		if len(shapes[shape]) != 1 {
			continue
		}

		to := shapes[shape][0]
		result = append(result, entry.New("", from, to))
		delete(shapes, shape)
		delete(before, key)
		delete(after, to.Key())
	}

	return result
}
