package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) SearchPages(
	_ context.Context,
	r server.SearchPagesRequestObject,
) (server.SearchPagesResponseObject, error) {
	result, e := s.confluence.Search(r.Params.Query)

	if e != nil {
		return server.SearchPages500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.SearchPages200JSONResponse(
		convert.ConfluencePages(result),
	), nil
}
