package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
)

func (s *Server) GetSearch(
	_ context.Context,
	r server.GetSearchRequestObject,
) (server.GetSearchResponseObject, error) {
	o := search_option.New(r.Params.Query, 10)

	if r.Params.Limit != nil {
		o.Limit = *r.Params.Limit
	}

	if r.Params.Collection != nil {
		o.Collection = *r.Params.Collection
	}

	o.Mode = "hybrid"

	if r.Params.Mode != nil {
		o.Mode = string(*r.Params.Mode)
	}

	o.Full = r.Params.Full != nil && *r.Params.Full

	if r.Params.Metadata != nil {
		o.Metadata = *r.Params.Metadata
	}

	if r.Params.SourceType != nil && *r.Params.SourceType != "" {
		if o.Metadata == nil {
			o.Metadata = map[string]string{}
		}

		o.Metadata[constant.SourceType] = *r.Params.SourceType
	}

	if r.Params.Exclude != nil {
		o.Exclude = *r.Params.Exclude
	}

	outcome := s.service.Search(o)
	result := server.GetSearch200JSONResponse{
		Results: convertSearchResults(outcome.Results),
	}

	if outcome.Degraded {
		result.Degraded = new(bool(true))
	}

	if len(outcome.Facets) > 0 {
		facets := convertFacets(outcome.Facets)
		result.Facets = &facets
	}

	return result, nil
}
