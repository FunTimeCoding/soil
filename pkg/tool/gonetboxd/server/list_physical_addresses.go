package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListPhysicalAddresses(
	_ context.Context,
	_ server.ListPhysicalAddressesRequestObject,
) (server.ListPhysicalAddressesResponseObject, error) {
	v, e := s.service.PhysicalAddresses()

	if forbidden.Is(e) {
		return server.ListPhysicalAddresses403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListPhysicalAddresses500JSONResponse(*s.captureDetail(e)),
			nil
	}

	return server.ListPhysicalAddresses200JSONResponse(
			convert.PhysicalAddressOwners(v),
		),
		nil
}
