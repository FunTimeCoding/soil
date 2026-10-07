package store_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"

func FindFacet(
	facets []result.Facet,
	key string,
) *result.Facet {
	for i := range facets {
		if facets[i].Key == key {
			return &facets[i]
		}
	}

	return nil
}
