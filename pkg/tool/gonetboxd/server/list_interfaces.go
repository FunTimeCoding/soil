package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListInterfaces(
	_ context.Context,
	r server.ListInterfacesRequestObject,
) (server.ListInterfacesResponseObject, error) {
	d, e := s.client.DeviceByName(r.Name)

	if e != nil {
		return server.ListInterfaces500JSONResponse(*s.captureDetail(e)), nil
	}

	interfaces, f := s.client.DeviceInterfaces(d.Identifier)

	if f != nil {
		return server.ListInterfaces500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.ListInterfaces200JSONResponse(
		convert.Interfaces(interfaces),
	), nil
}
