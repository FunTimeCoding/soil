package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) DeriveHardwareAddress(
	_ context.Context,
	r server.DeriveHardwareAddressRequestObject,
) (server.DeriveHardwareAddressResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.DeriveHardwareAddress404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.DeriveHardwareAddress400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DeriveHardwareAddress500JSONResponse(*s.answer(e)), nil
	}

	address, holder, g := s.service.DeriveHardwareAddress(
		instance,
		r.Params.Identifier,
	)

	switch {
	case validation.Is(g):
		return server.DeriveHardwareAddress400JSONResponse(*clientError(g)), nil
	case not_found.Is(g):
		return server.DeriveHardwareAddress404JSONResponse(*clientError(g)), nil
	case forbidden.Is(g):
		return server.DeriveHardwareAddress403JSONResponse(*clientError(g)), nil
	case unavailable.Is(g):
		return server.DeriveHardwareAddress503JSONResponse(*clientError(g)), nil
	case g != nil:
		return server.DeriveHardwareAddress500JSONResponse(*s.answer(g)), nil
	}

	return server.DeriveHardwareAddress200JSONResponse{
		Instance:        instance,
		Identifier:      r.Params.Identifier,
		HardwareAddress: address,
		InUseBy:         holder,
	}, nil
}
