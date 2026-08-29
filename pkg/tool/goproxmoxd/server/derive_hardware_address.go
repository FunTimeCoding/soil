package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) DeriveHardwareAddress(
	_ context.Context,
	r server.DeriveHardwareAddressRequestObject,
) (server.DeriveHardwareAddressResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.DeriveHardwareAddress400JSONResponse(*clientError(e)), nil
	}

	address, holder, g := s.service.DeriveHardwareAddress(
		instance,
		r.Params.Identifier,
	)

	if g != nil {
		if validation.Is(g) {
			return server.DeriveHardwareAddress400JSONResponse(
				*clientError(g),
			), nil
		}

		return server.DeriveHardwareAddress500JSONResponse(
			*s.captureDetail(g),
		), nil
	}

	return server.DeriveHardwareAddress200JSONResponse{
		Instance:        instance,
		Identifier:      r.Params.Identifier,
		HardwareAddress: address,
		InUseBy:         holder,
	}, nil
}
