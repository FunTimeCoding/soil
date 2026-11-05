package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListVirtualAddresses(
	x context.Context,
	r server.ListVirtualAddressesRequestObject,
) (server.ListVirtualAddressesResponseObject, error) {
	addresses, e := s.service.VirtualMachineAddresses(r.Name)

	if not_found.Is(e) {
		return server.ListVirtualAddresses404JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.ListVirtualAddresses503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListVirtualAddresses500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.ListVirtualAddresses200JSONResponse(
			convert.Addresses(addresses),
		),
		nil
}
