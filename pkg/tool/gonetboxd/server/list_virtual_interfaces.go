package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListVirtualInterfaces(
	_ context.Context,
	r server.ListVirtualInterfacesRequestObject,
) (server.ListVirtualInterfacesResponseObject, error) {
	interfaces, e := s.service.VirtualInterfaces(r.Name)

	if e != nil {
		return server.ListVirtualInterfaces500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.ListVirtualInterfaces200JSONResponse(
		convert.Interfaces(interfaces),
	), nil
}
