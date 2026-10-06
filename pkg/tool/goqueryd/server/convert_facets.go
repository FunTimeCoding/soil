package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"
)

func convertFacets(facets []search.Facet) []server.Facet {
	converted := make([]server.Facet, len(facets))

	for i, f := range facets {
		converted[i] = *convertFacet(f)
	}

	return converted
}
