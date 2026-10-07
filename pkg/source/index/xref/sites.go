package xref

import "github.com/funtimecoding/soil/pkg/source/index/record"

func (i *Index) Sites(target string) map[string][]*record.Site {
	result := make(map[string][]*record.Site)

	for unit, r := range i.units {
		if sites := r.Targets[target]; len(sites) > 0 {
			result[unit] = sites
		}
	}

	return result
}
