package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateVirtualInterface(
	_ context.Context,
	r server.CreateVirtualInterfaceRequestObject,
) (server.CreateVirtualInterfaceResponseObject, error) {
	i, e := s.service.CreateVirtualInterface(r.Name, r.Body.Name)

	if not_found.Is(e) {
		return server.CreateVirtualInterface404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateVirtualInterface500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateVirtualInterface201JSONResponse(
		*convert.VirtualInterface(i),
	), nil
}
