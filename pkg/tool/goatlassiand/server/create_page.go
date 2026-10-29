package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) CreatePage(
	_ context.Context,
	r server.CreatePageRequestObject,
) (server.CreatePageResponseObject, error) {
	result, e := s.service.CreatePage(
		r.Body.SpaceIdentifier,
		r.Body.ParentIdentifier,
		r.Body.Title,
		r.Body.Body,
		false,
	)

	if forbidden.Is(e) {
		return server.CreatePage403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.CreatePage404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.CreatePage503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreatePage500JSONResponse(*s.answer(e)), nil
	}

	return server.CreatePage201JSONResponse(
		*convert.ConfluencePageDetail(result),
	), nil
}
