package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateDeviceRole(
	x context.Context,
	r server.CreateDeviceRoleRequestObject,
) (server.CreateDeviceRoleResponseObject, error) {
	var virtualMachine bool

	if r.Body.VirtualMachine != nil {
		virtualMachine = *r.Body.VirtualMachine
	}

	role, e := s.service.CreateDeviceRole(r.Body.Name, virtualMachine)

	if validation.Is(e) {
		return server.CreateDeviceRole400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreateDeviceRole503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateDeviceRole500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.CreateDeviceRole201JSONResponse(*convert.DeviceRole(role)), nil
}
