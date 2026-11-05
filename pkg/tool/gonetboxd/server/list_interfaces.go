package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListInterfaces(
	x context.Context,
	r server.ListInterfacesRequestObject,
) (server.ListInterfacesResponseObject, error) {
	interfaces, e := s.service.DeviceInterfaces(r.Name)

	if not_found.Is(e) {
		return server.ListInterfaces404JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.ListInterfaces503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListInterfaces500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListInterfaces200JSONResponse(convert.Interfaces(interfaces)), nil
}
