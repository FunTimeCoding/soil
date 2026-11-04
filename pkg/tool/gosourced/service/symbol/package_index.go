package symbol

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

func packageIndex(all []*packages.Package) map[string]*types.Package {
	result := map[string]*types.Package{}
	var queue []*types.Package

	for _, p := range all {
		if p.Types != nil {
			queue = append(queue, p.Types)
		}
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if _, seen := result[p.Path()]; seen {
			continue
		}

		result[p.Path()] = p
		queue = append(queue, p.Imports()...)
	}

	return result
}
