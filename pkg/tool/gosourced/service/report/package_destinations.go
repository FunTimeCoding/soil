package report

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
	"slices"
)

func packageDestinations(pairs []*entry.Entry) map[string][]string {
	result := map[string][]string{}

	for _, p := range pairs {
		if !slices.Contains(result[p.From.Path], p.To.Path) {
			result[p.From.Path] = append(result[p.From.Path], p.To.Path)
		}
	}

	return result
}
