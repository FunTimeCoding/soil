package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateVirtualTunnelTermination(
	_ context.Context,
	r server.CreateVirtualTunnelTerminationRequestObject,
) (server.CreateVirtualTunnelTerminationResponseObject, error) {
	t, e := s.client.TunnelByName(r.Body.Tunnel)

	if e != nil {
		return server.CreateVirtualTunnelTermination500JSONResponse(*s.captureDetail(e)), nil
	}

	vm, f := s.client.VirtualMachineByName(r.Name)

	if f != nil {
		return server.CreateVirtualTunnelTermination500JSONResponse(*s.captureDetail(f)), nil
	}

	i, g := s.client.VirtualMachineInterfaceByName(vm, r.Body.Interface)

	if g != nil {
		return server.CreateVirtualTunnelTermination500JSONResponse(*s.captureDetail(g)), nil
	}

	tt, h := s.client.CreateTunnelTermination(
		t,
		constant.VirtualInterfaceAddress,
		int64(i.GetId()),
		r.Body.Role,
	)

	if h != nil {
		return server.CreateVirtualTunnelTermination500JSONResponse(*s.captureDetail(h)), nil
	}

	return server.CreateVirtualTunnelTermination201JSONResponse(
		*convert.TunnelTermination(tt),
	), nil
}
