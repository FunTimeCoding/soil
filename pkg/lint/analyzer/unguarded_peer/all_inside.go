package unguarded_peer

import "go/types"

func allInside(
	l []*Reference,
	inside map[types.Object]bool,
) bool {
	for _, r := range l {
		if !r.Guarded && !inside[r.Enclosing] {
			return false
		}
	}

	return len(l) > 0
}
