package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListVirtualInterfaces(
	_ context.Context,
	r server.ListVirtualInterfacesRequestObject,
) (server.ListVirtualInterfacesResponseObject, error) {
	interfaces, e := s.service.VirtualInterfaces(r.Name)

	if not_found.Is(e) {
		return server.ListVirtualInterfaces404JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.ListVirtualInterfaces403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListVirtualInterfaces500JSONResponse(*s.captureDetail(e)),
			nil
	}

	return server.ListVirtualInterfaces200JSONResponse(
			convert.Interfaces(interfaces),
		),
		nil
}
