package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListVirtualAddresses(
	_ context.Context,
	r server.ListVirtualAddressesRequestObject,
) (server.ListVirtualAddressesResponseObject, error) {
	addresses, e := s.client.VirtualMachineAddresses(r.Name)

	if e != nil {
		return server.ListVirtualAddresses500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListVirtualAddresses200JSONResponse(
		convert.Addresses(addresses),
	), nil
}
