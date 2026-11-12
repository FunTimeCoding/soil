package census

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/census/reading"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
	"slices"
)

func callerResults(callers map[reading.CallerKey]int) []*result.CensusCaller {
	entries := make([]*result.CensusCaller, 0, len(callers))

	for k, count := range callers {
		entries = append(
			entries,
			result.NewCensusCaller(k.Kind, k.PackagePath, k.Function, count),
		)
	}

	slices.SortFunc(
		entries,
		func(a, b *result.CensusCaller) int {
			return cmp.Or(
				cmp.Compare(a.Kind, b.Kind),
				cmp.Compare(a.Package, b.Package),
				cmp.Compare(a.Function, b.Function),
			)
		},
	)

	return entries
}
