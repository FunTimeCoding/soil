package usage

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
	"maps"
	"slices"
)

func unloadedResults(missing unloaded) []*result.UsageUnloaded {
	r := []*result.UsageUnloaded{}

	for _, k := range slices.Sorted(maps.Keys(missing)) {
		r = append(r, missing[k])
	}

	return r
}
