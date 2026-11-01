package report

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
	"slices"
)

func sortEntries(entries []*entry.Entry) {
	slices.SortFunc(
		entries,
		func(
			a *entry.Entry,
			b *entry.Entry,
		) int {
			return cmp.Or(
				cmp.Compare(
					slices.Index(constant.Classes, a.Class),
					slices.Index(constant.Classes, b.Class),
				),
				cmp.Compare(a.Key(), b.Key()),
			)
		},
	)
}
