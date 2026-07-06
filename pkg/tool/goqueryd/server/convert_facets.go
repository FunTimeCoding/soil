package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func convertFacets(facets []store.Facet) []server.Facet {
	converted := make([]server.Facet, len(facets))

	for i, f := range facets {
		converted[i] = *convertFacet(f)
	}

	return converted
}
