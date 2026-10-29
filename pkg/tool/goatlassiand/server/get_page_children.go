package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetPageChildren(
	_ context.Context,
	r server.GetPageChildrenRequestObject,
) (server.GetPageChildrenResponseObject, error) {
	result, e := s.service.PageChildren(r.Identifier)

	if forbidden.Is(e) {
		return server.GetPageChildren403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.GetPageChildren404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetPageChildren503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetPageChildren500JSONResponse(*s.answer(e)), nil
	}

	return server.GetPageChildren200JSONResponse(
		convert.ConfluencePagesFromPages(result),
	), nil
}
