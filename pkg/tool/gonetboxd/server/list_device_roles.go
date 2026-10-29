package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListDeviceRoles(
	_ context.Context,
	_ server.ListDeviceRolesRequestObject,
) (server.ListDeviceRolesResponseObject, error) {
	roles, e := s.service.DeviceRoles()

	if forbidden.Is(e) {
		return server.ListDeviceRoles403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListDeviceRoles500JSONResponse(*s.captureDetail(e)), nil
	}

	result := make([]*server.DeviceRole, 0, len(roles))

	for _, r := range roles {
		result = append(result, convert.DeviceRole(r))
	}

	return server.ListDeviceRoles200JSONResponse(result), nil
}
