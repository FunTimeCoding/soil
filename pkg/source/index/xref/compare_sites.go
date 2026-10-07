package xref

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/source/index/record"
)

func compareSites(
	a *record.Site,
	b *record.Site,
) int {
	return cmp.Or(
		cmp.Compare(a.File, b.File),
		cmp.Compare(a.Line, b.Line),
		cmp.Compare(a.Column, b.Column),
	)
}
