package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateAddress(
	_ context.Context,
	r server.CreateAddressRequestObject,
) (server.CreateAddressResponseObject, error) {
	status := ""

	if r.Body.Status != nil {
		status = *r.Body.Status
	}

	a, e := s.service.CreateAddress(
		r.Name,
		r.Body.Interface,
		r.Body.Address,
		status,
	)

	if not_found.Is(e) {
		return server.CreateAddress404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.CreateAddress400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.CreateAddress403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateAddress500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateAddress201JSONResponse{
		Identifier: a.Identifier,
		Address:    a.Address.String(),
	}, nil
}
