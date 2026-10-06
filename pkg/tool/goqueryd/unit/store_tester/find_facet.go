package store_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"

func FindFacet(
	facets []search.Facet,
	key string,
) *search.Facet {
	for i := range facets {
		if facets[i].Key == key {
			return &facets[i]
		}
	}

	return nil
}
