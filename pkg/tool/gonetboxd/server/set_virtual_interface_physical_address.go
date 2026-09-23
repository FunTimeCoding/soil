package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) SetVirtualInterfacePhysicalAddress(
	_ context.Context,
	r server.SetVirtualInterfacePhysicalAddressRequestObject,
) (server.SetVirtualInterfacePhysicalAddressResponseObject, error) {
	m, e := s.client.VirtualMachineByName(r.Name)

	if e != nil {
		return server.SetVirtualInterfacePhysicalAddress500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	updated, f := s.client.UpdateVirtualInterface(
		m,
		r.Interface,
		network.PhysicalAddress(r.Body.Address),
	)

	if f != nil {
		return server.SetVirtualInterfacePhysicalAddress500JSONResponse(
			*s.captureDetail(f),
		), nil
	}

	return server.SetVirtualInterfacePhysicalAddress200JSONResponse(
		*convert.VirtualInterfacePhysicalAddress(updated, r.Body.Address),
	), nil
}
