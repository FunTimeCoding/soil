package unguarded_peer

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

func insideFunctions(
	p *packages.Package,
	wrappers map[types.Object]bool,
) map[types.Object]bool {
	references := functionReferences(p, wrappers)
	result := make(map[types.Object]bool)

	for changed := true; changed; {
		changed = false

		for o, l := range references {
			if result[o] || !allInside(l, result) {
				continue
			}

			result[o] = true
			changed = true
		}
	}

	return result
}
