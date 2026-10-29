package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) SearchPages(
	_ context.Context,
	r server.SearchPagesRequestObject,
) (server.SearchPagesResponseObject, error) {
	result, e := s.service.SearchPages(r.Params.Query)

	if forbidden.Is(e) {
		return server.SearchPages403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.SearchPages503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.SearchPages500JSONResponse(*s.answer(e)), nil
	}

	return server.SearchPages200JSONResponse(convert.ConfluencePages(result)), nil
}
