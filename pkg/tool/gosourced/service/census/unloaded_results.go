package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
	"maps"
	"slices"
)

func unloadedResults(missing unloaded) []*result.CensusUnloaded {
	r := []*result.CensusUnloaded{}

	for _, k := range slices.Sorted(maps.Keys(missing)) {
		r = append(r, missing[k])
	}

	return r
}
