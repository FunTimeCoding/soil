package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateTag(
	x context.Context,
	r server.CreateTagRequestObject,
) (server.CreateTagResponseObject, error) {
	t, e := s.service.CreateTag(r.Body.Name)

	if validation.Is(e) {
		return server.CreateTag400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreateTag503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateTag500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.CreateTag201JSONResponse{
		Identifier: t.Identifier,
		Name:       t.Name,
	}, nil
}
