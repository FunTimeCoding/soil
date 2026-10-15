package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListVirtualAddresses(
	_ context.Context,
	r server.ListVirtualAddressesRequestObject,
) (server.ListVirtualAddressesResponseObject, error) {
	addresses, e := s.client.VirtualMachineAddresses(r.Name)

	if not_found.Is(e) {
		return server.ListVirtualAddresses404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListVirtualAddresses500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListVirtualAddresses200JSONResponse(
		convert.Addresses(addresses),
	), nil
}
