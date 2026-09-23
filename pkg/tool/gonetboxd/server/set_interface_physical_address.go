package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) SetInterfacePhysicalAddress(
	_ context.Context,
	r server.SetInterfacePhysicalAddressRequestObject,
) (server.SetInterfacePhysicalAddressResponseObject, error) {
	d, e := s.client.DeviceByName(r.Name)

	if e != nil {
		return server.SetInterfacePhysicalAddress500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	i, f := s.client.DeviceInterfaceByName(d, r.Interface)

	if f != nil {
		return server.SetInterfacePhysicalAddress500JSONResponse(
			*s.captureDetail(f),
		), nil
	}

	updated, g := s.client.UpdateInterface(
		d,
		r.Interface,
		i.Type,
		network.PhysicalAddress(r.Body.Address),
	)

	if g != nil {
		return server.SetInterfacePhysicalAddress500JSONResponse(
			*s.captureDetail(g),
		), nil
	}

	return server.SetInterfacePhysicalAddress200JSONResponse(
		*convert.InterfacePhysicalAddress(updated),
	), nil
}
