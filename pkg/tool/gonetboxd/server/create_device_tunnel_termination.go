package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateDeviceTunnelTermination(
	_ context.Context,
	r server.CreateDeviceTunnelTerminationRequestObject,
) (server.CreateDeviceTunnelTerminationResponseObject, error) {
	t, e := s.client.TunnelByName(r.Body.Tunnel)

	if e != nil {
		return server.CreateDeviceTunnelTermination500JSONResponse(*s.captureDetail(e)), nil
	}

	d, f := s.client.DeviceByName(r.Name)

	if f != nil {
		return server.CreateDeviceTunnelTermination500JSONResponse(*s.captureDetail(f)), nil
	}

	i, g := s.client.DeviceInterfaceByName(d, r.Body.Interface)

	if g != nil {
		return server.CreateDeviceTunnelTermination500JSONResponse(*s.captureDetail(g)), nil
	}

	tt, h := s.client.CreateTunnelTermination(
		t,
		constant.InterfaceAddress,
		int64(i.Identifier),
		r.Body.Role,
	)

	if h != nil {
		return server.CreateDeviceTunnelTermination500JSONResponse(*s.captureDetail(h)), nil
	}

	return server.CreateDeviceTunnelTermination201JSONResponse(
		*convert.TunnelTermination(tt),
	), nil
}
