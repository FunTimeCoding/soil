package service

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"slices"
)

func tagCounts(counts map[string]int) []*record.TagCount {
	result := make([]*record.TagCount, 0, len(counts))

	for tag, count := range counts {
		result = append(result, record.NewTagCount(tag, count))
	}

	slices.SortFunc(
		result,
		func(
			a *record.TagCount,
			b *record.TagCount,
		) int {
			return cmp.Or(
				cmp.Compare(b.Count, a.Count),
				cmp.Compare(a.Tag, b.Tag),
			)
		},
	)

	return result
}
