package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateDeviceRole(
	_ context.Context,
	r server.CreateDeviceRoleRequestObject,
) (server.CreateDeviceRoleResponseObject, error) {
	var virtualMachine bool

	if r.Body.VirtualMachine != nil {
		virtualMachine = *r.Body.VirtualMachine
	}

	role, e := s.client.CreateDeviceRole(r.Body.Name, virtualMachine)

	if e != nil {
		return server.CreateDeviceRole500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateDeviceRole201JSONResponse(*convert.DeviceRole(role)), nil
}
