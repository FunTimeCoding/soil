package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) GetMetadata(
	_ context.Context,
	r server.GetMetadataRequestObject,
) (server.GetMetadataResponseObject, error) {
	key := ""

	if r.Params.Key != nil {
		key = *r.Params.Key
	}

	if key != "" {
		facets := s.service.CollectionFacetsForKey(r.Params.Collection, key)
		result := make(server.GetMetadata200JSONResponse, len(facets))

		for i, f := range facets {
			result[i] = convertFacet(f)
		}

		return result, nil
	}

	facets := s.service.CollectionFacets(r.Params.Collection, nil)
	result := make(server.GetMetadata200JSONResponse, len(facets))

	for i, f := range facets {
		result[i] = convertFacet(f)
	}

	return result, nil
}
