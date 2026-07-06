package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	queryConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) GetList(
	_ context.Context,
	r server.GetListRequestObject,
) (server.GetListResponseObject, error) {
	var metadata map[string]string

	if r.Params.Metadata != nil {
		metadata = *r.Params.Metadata
	}

	if r.Params.SourceType != nil && *r.Params.SourceType != "" {
		if metadata == nil {
			metadata = map[string]string{}
		}

		metadata[queryConstant.SourceType] = *r.Params.SourceType
	}

	limit := 10

	if r.Params.Limit != nil {
		limit = *r.Params.Limit
	}

	offset := 0

	if r.Params.Offset != nil {
		offset = *r.Params.Offset
	}

	outcome, e := s.service.ListDocuments(
		r.Params.Collection,
		metadata,
		limit,
		offset,
		r.Params.Full != nil && *r.Params.Full,
	)

	if e != nil {
		return server.GetList500JSONResponse(
			*s.captureFail(
				e,
				constant.UnexpectedError,
			),
		), nil
	}

	result := server.GetList200JSONResponse{
		Results: convertSearchResults(outcome.Results),
	}

	if len(outcome.Facets) > 0 {
		facets := convertFacets(outcome.Facets)
		result.Facets = &facets
	}

	return result, nil
}
