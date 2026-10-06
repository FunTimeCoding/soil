package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern_site"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
	"sort"
)

func groupEntries(entries []*pattern_site.Entry) []*result.Group {
	sort.Slice(
		entries,
		func(i, j int) bool {
			if entries[i].Location.File != entries[j].Location.File {
				return entries[i].Location.File < entries[j].Location.File
			}

			return entries[i].Location.Line < entries[j].Location.Line
		},
	)
	groups := map[string]*result.Group{}
	var order []*result.Group

	for _, entry := range entries {
		group, exists := groups[entry.Shape]

		if !exists {
			group = result.NewGroup(entry.Shape, entry.Exemplar, nil)
			groups[entry.Shape] = group
			order = append(order, group)
		}

		group.Locations = append(group.Locations, entry.Location)
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
