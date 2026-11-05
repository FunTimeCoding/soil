package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListPhysicalAddresses(
	x context.Context,
	_ server.ListPhysicalAddressesRequestObject,
) (server.ListPhysicalAddressesResponseObject, error) {
	v, e := s.service.PhysicalAddresses()

	if unavailable.Is(e) {
		return server.ListPhysicalAddresses503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListPhysicalAddresses500JSONResponse(
				*s.captureDetail(x, e),
			),
			nil
	}

	return server.ListPhysicalAddresses200JSONResponse(
			convert.PhysicalAddressOwners(v),
		),
		nil
}
