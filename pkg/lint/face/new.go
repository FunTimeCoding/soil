package face

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

func New(loaded []*packages.Package) *Set {
	result := &Set{byMethod: make(map[string][]*types.Interface)}
	seen := make(map[string]bool)

	for _, p := range loaded {
		if p.Types == nil {
			continue
		}

		result.collect(p.Types, seen)
	}

	return result
}
