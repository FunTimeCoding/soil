package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) UpdatePage(
	_ context.Context,
	r server.UpdatePageRequestObject,
) (server.UpdatePageResponseObject, error) {
	message := ""

	if r.Body.Message != nil {
		message = *r.Body.Message
	}

	result, e := s.service.UpdatePage(
		r.Identifier,
		r.Body.Title,
		r.Body.Body,
		message,
	)

	if forbidden.Is(e) {
		return server.UpdatePage403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.UpdatePage404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.UpdatePage503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.UpdatePage500JSONResponse(*s.answer(e)), nil
	}

	return server.UpdatePage200JSONResponse(
		*convert.ConfluencePageDetail(result),
	), nil
}
