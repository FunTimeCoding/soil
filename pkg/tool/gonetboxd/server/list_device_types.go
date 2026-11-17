package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListDeviceTypes(
	x context.Context,
	_ server.ListDeviceTypesRequestObject,
) (server.ListDeviceTypesResponseObject, error) {
	types, e := s.service.DeviceTypes()

	if unavailable.Is(e) {
		return server.ListDeviceTypes503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListDeviceTypes500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.ListDeviceTypes200JSONResponse(convert.DeviceTypes(types)), nil
}
