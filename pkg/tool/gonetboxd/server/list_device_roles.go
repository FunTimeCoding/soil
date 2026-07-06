package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListDeviceRoles(
	_ context.Context,
	_ server.ListDeviceRolesRequestObject,
) (server.ListDeviceRolesResponseObject, error) {
	roles, e := s.client.DeviceRoles()

	if e != nil {
		return server.ListDeviceRoles500JSONResponse(*s.captureDetail(e)), nil
	}

	result := make([]*server.DeviceRole, 0, len(roles))

	for _, r := range roles {
		result = append(
			result,
			&server.DeviceRole{Identifier: r.Identifier, Name: r.Name},
		)
	}

	return server.ListDeviceRoles200JSONResponse(result), nil
}
