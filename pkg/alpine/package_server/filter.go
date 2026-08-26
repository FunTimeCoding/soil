package package_server

import "github.com/funtimecoding/soil/pkg/alpine/index"

func Filter(
	listings []*Listing,
	name string,
) []*Listing {
	if name == "" {
		return listings
	}

	var result []*Listing

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
