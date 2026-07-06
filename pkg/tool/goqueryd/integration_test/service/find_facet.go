package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store"

func findFacet(
	facets []store.Facet,
	key string,
) *store.Facet {
	for _, f := range facets {
		if f.Key == key {
			return &f
		}
	}

	return nil
}
