package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) PostIndex(
	_ context.Context,
	r server.PostIndexRequestObject,
) (server.PostIndexResponseObject, error) {
	collection := ""

	if r.Body.Collection != nil {
		collection = *r.Body.Collection
	}

	results, e := s.service.IndexCollections(collection)

	if e != nil {
		return server.PostIndex500JSONResponse(
			*s.captureFail(
				e,
				constant.UnexpectedError,
			),
		), nil
	}

	converted := make(server.PostIndex200JSONResponse, len(results))

	for i, r := range results {
		converted[i] = &server.IndexResult{
			Collection: r.Collection,
			Indexed:    r.Indexed,
			Updated:    r.Updated,
			Unchanged:  r.Unchanged,
			Removed:    r.Removed,
		}
	}

	return converted, nil
}
