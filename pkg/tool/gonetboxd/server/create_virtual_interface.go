package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateVirtualInterface(
	x context.Context,
	r server.CreateVirtualInterfaceRequestObject,
) (server.CreateVirtualInterfaceResponseObject, error) {
	i, e := s.service.CreateVirtualInterface(r.Name, r.Body.Name)

	if not_found.Is(e) {
		return server.CreateVirtualInterface404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.CreateVirtualInterface400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreateVirtualInterface503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateVirtualInterface500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.CreateVirtualInterface201JSONResponse(
			*convert.VirtualInterface(i),
		),
		nil
}
