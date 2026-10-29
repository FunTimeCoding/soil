package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetPage(
	_ context.Context,
	r server.GetPageRequestObject,
) (server.GetPageResponseObject, error) {
	result, e := s.service.Page(r.Identifier, false)

	if forbidden.Is(e) {
		return server.GetPage403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.GetPage404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetPage503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetPage500JSONResponse(*s.answer(e)), nil
	}

	return server.GetPage200JSONResponse(*convert.ConfluencePageDetail(result)), nil
}
