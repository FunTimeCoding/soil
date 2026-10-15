package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) SetVirtualInterfacePhysicalAddress(
	_ context.Context,
	r server.SetVirtualInterfacePhysicalAddressRequestObject,
) (server.SetVirtualInterfacePhysicalAddressResponseObject, error) {
	updated, e := s.service.SetVirtualInterfacePhysicalAddress(
		r.Name,
		r.Interface,
		r.Body.Address,
	)

	if not_found.Is(e) {
		return server.SetVirtualInterfacePhysicalAddress404JSONResponse(
			*clientError(e),
		), nil
	}

	if e != nil {
		return server.SetVirtualInterfacePhysicalAddress500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.SetVirtualInterfacePhysicalAddress200JSONResponse(
		*convert.VirtualInterfacePhysicalAddress(updated, r.Body.Address),
	), nil
}
