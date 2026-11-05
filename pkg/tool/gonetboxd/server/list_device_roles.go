package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListDeviceRoles(
	x context.Context,
	_ server.ListDeviceRolesRequestObject,
) (server.ListDeviceRolesResponseObject, error) {
	roles, e := s.service.DeviceRoles()

	if unavailable.Is(e) {
		return server.ListDeviceRoles503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListDeviceRoles500JSONResponse(*s.captureDetail(x, e)), nil
	}

	result := make([]*server.DeviceRole, 0, len(roles))

	for _, r := range roles {
		result = append(result, convert.DeviceRole(r))
	}

	return server.ListDeviceRoles200JSONResponse(result), nil
}
