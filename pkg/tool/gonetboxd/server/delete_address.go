package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) DeleteAddress(
	_ context.Context,
	r server.DeleteAddressRequestObject,
) (server.DeleteAddressResponseObject, error) {
	e := s.service.DeleteInternet(r.Identifier)

	if validation.Is(e) {
		return server.DeleteAddress400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.DeleteAddress403JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.DeleteAddress409JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DeleteAddress500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteAddress204Response{}, nil
}
