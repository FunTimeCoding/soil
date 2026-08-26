package convert

import (
	"github.com/funtimecoding/soil/pkg/alpine/package_server"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/generated/server"
)

func Listings(listings []*package_server.Listing) []server.Listing {
	result := []server.Listing{}

	for _, l := range listings {
		entries := []server.Entry{}

		for _, entry := range l.Packages {
			entries = append(
				entries,
				server.Entry{
					Name:         entry.Name,
					Version:      entry.Version,
					Architecture: entry.Architecture,
				},
			)
		}

		result = append(
			result,
			server.Listing{
				Version:      l.Version,
				Repository:   l.Repository,
				Architecture: l.Architecture,
				Packages:     entries,
			},
		)
	}

	return result
}
