package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"

func findFacet(
	facets []result.Facet,
	key string,
) *result.Facet {
	for _, f := range facets {
		if f.Key == key {
			return &f
		}
	}

	return nil
}
