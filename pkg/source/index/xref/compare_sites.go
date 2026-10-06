package xref

import "cmp"

func compareSites(
	a *Site,
	b *Site,
) int {
	return cmp.Or(
		cmp.Compare(a.File, b.File),
		cmp.Compare(a.Line, b.Line),
		cmp.Compare(a.Column, b.Column),
	)
}
