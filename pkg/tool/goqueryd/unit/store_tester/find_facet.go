package store_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store"

func FindFacet(
	facets []store.Facet,
	key string,
) *store.Facet {
	for i := range facets {
		if facets[i].Key == key {
			return &facets[i]
		}
	}

	return nil
}
