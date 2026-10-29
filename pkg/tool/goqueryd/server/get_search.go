package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
)

func (s *Server) GetSearch(
	x context.Context,
	r server.GetSearchRequestObject,
) (server.GetSearchResponseObject, error) {
	o := search_option.New(r.Params.Query, 10)

	if r.Params.Limit != nil {
		o.Limit = *r.Params.Limit
	}

	if r.Params.Collection != nil {
		o.Collection = *r.Params.Collection
	}

	o.Mode = constant.HybridMode

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

	outcome, e := s.service.Search(x, o)

	if unavailable.Is(e) {
		return server.GetSearch503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetSearch500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	result := server.GetSearch200JSONResponse{
		Results: convertSearchResults(outcome.Results),
	}

	if len(outcome.Facets) > 0 {
		facets := convertFacets(outcome.Facets)
		result.Facets = &facets
	}

	return result, nil
}
