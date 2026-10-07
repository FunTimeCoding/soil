package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
)

func convertFacets(facets []result.Facet) []server.Facet {
	converted := make([]server.Facet, len(facets))

	for i, f := range facets {
		converted[i] = *convertFacet(f)
	}

	return converted
}
