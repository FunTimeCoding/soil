package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) UpdateDeviceRole(
	x context.Context,
	r server.UpdateDeviceRoleRequestObject,
) (server.UpdateDeviceRoleResponseObject, error) {
	role, e := s.service.UpdateDeviceRole(r.Name, r.Body.Name)

	if not_found.Is(e) {
		return server.UpdateDeviceRole404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.UpdateDeviceRole400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.UpdateDeviceRole503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.UpdateDeviceRole500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.UpdateDeviceRole200JSONResponse(*convert.DeviceRole(role)), nil
}
