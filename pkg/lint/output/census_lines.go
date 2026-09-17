package output

import (
	"cmp"
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"maps"
	"slices"
)

func CensusLines(entries []*Unchecked) []string {
	sorted := slices.Clone(entries)
	slices.SortFunc(
		sorted,
		func(
			x *Unchecked,
			y *Unchecked,
		) int {
			return cmp.Or(
				cmp.Compare(x.Reason, y.Reason),
				cmp.Compare(x.Path, y.Path),
				cmp.Compare(x.Line, y.Line),
			)
		},
	)
	var result []string
	counts := map[constant.Reason]int{}

	for _, u := range sorted {
		result = append(
			result,
			fmt.Sprintf(
				"%s:%d: unchecked %s %s",
				u.Path,
				u.Line,
				u.Reason,
				u.Span,
			),
		)
		counts[u.Reason]++
	}

	for _, reason := range slices.Sorted(maps.Keys(counts)) {
		result = append(
			result,
			fmt.Sprintf("unchecked %s: %d", reason, counts[reason]),
		)
	}

	if len(sorted) > 0 {
		result = append(result, fmt.Sprintf("unchecked: %d", len(sorted)))
	}

	return result
}
