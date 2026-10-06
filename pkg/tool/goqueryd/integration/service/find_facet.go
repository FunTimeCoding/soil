package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"

func findFacet(
	facets []search.Facet,
	key string,
) *search.Facet {
	for _, f := range facets {
		if f.Key == key {
			return &f
		}
	}

	return nil
}
