package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) DeleteAddress(
	x context.Context,
	r server.DeleteAddressRequestObject,
) (server.DeleteAddressResponseObject, error) {
	e := s.service.DeleteInternet(r.Identifier)

	if validation.Is(e) {
		return server.DeleteAddress400JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.DeleteAddress409JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.DeleteAddress503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteAddress500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.DeleteAddress204Response{}, nil
}
