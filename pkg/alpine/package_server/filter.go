package package_server

import (
	"github.com/funtimecoding/soil/pkg/alpine/index"
	"github.com/funtimecoding/soil/pkg/alpine/types/listing"
)

func Filter(
	listings []*listing.Listing,
	name string,
) []*listing.Listing {
	if name == "" {
		return listings
	}

	var result []*listing.Listing

	for _, l := range listings {
		var entries []*index.Entry

		for _, entry := range l.Packages {
			if entry.Name == name {
				entries = append(entries, entry)
			}
		}

		if len(entries) == 0 {
			continue
		}

		filtered := *l
		filtered.Packages = entries
		result = append(result, &filtered)
	}

	return result
}
