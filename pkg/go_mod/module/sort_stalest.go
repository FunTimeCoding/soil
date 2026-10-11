package module

import "slices"

func SortStalest(v []*Module) {
	slices.SortStableFunc(
		v,
		func(
			a *Module,
			b *Module,
		) int {
			switch {
			case a.Release == nil && b.Release == nil:
				return 0
			case a.Release == nil:
				return 1
			case b.Release == nil:
				return -1
			}

			return a.Release.Compare(*b.Release)
		},
	)
}
