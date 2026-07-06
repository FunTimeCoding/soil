package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateDeviceRole(
	_ context.Context,
	r server.CreateDeviceRoleRequestObject,
) (server.CreateDeviceRoleResponseObject, error) {
	role, e := s.client.CreateDeviceRole(r.Body.Name)

	if e != nil {
		return server.CreateDeviceRole500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateDeviceRole201JSONResponse{
		Identifier: role.Identifier, Name: role.Name,
	}, nil
}
