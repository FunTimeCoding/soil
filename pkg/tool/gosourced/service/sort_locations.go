package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result/location"
	"sort"
)

func sortLocations(locations []*location.Location) []*location.Location {
	sort.SliceStable(
		locations,
		func(
			i int,
			j int,
		) bool {
			if locations[i].File != locations[j].File {
				return locations[i].File < locations[j].File
			}

			return locations[i].Line < locations[j].Line
		},
	)

	return locations
}
