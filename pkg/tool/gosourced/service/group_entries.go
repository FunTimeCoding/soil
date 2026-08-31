package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
	"sort"
)

func groupEntries(entries []*siteEntry) []*result.Group {
	sort.Slice(
		entries,
		func(i, j int) bool {
			if entries[i].location.File != entries[j].location.File {
				return entries[i].location.File < entries[j].location.File
			}

			return entries[i].location.Line < entries[j].location.Line
		},
	)
	groups := map[string]*result.Group{}
	var order []*result.Group

	for _, entry := range entries {
		group, exists := groups[entry.shape]

		if !exists {
			group = result.NewGroup(entry.shape, entry.exemplar, nil)
			groups[entry.shape] = group
			order = append(order, group)
		}

		group.Locations = append(group.Locations, entry.location)
	}

	sort.Slice(
		order,
		func(i, j int) bool {
			if len(order[i].Locations) != len(order[j].Locations) {
				return len(order[i].Locations) > len(order[j].Locations)
			}

			return order[i].Shape < order[j].Shape
		},
	)

	return order
}
