package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) DeleteCable(
	x context.Context,
	r server.DeleteCableRequestObject,
) (server.DeleteCableResponseObject, error) {
	e := s.service.DeleteCable(r.Identifier)

	if not_found.Is(e) {
		return server.DeleteCable404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.DeleteCable400JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.DeleteCable409JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.DeleteCable503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteCable500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.DeleteCable204Response{}, nil
}
