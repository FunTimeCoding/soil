package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) ListPages(
	_ context.Context,
	r server.ListPagesRequestObject,
) (server.ListPagesResponseObject, error) {
	status := ""

	if r.Params.Status != nil {
		status = *r.Params.Status
	}

	result, e := s.service.Pages(r.Params.Space, status)

	if forbidden.Is(e) {
		return server.ListPages403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.ListPages404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.ListPages503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListPages500JSONResponse(*s.answer(e)), nil
	}

	return server.ListPages200JSONResponse(
		convert.ConfluencePagesFromPages(result),
	), nil
}
