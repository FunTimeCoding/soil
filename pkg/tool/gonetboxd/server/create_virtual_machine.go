package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateVirtualMachine(
	x context.Context,
	r server.CreateVirtualMachineRequestObject,
) (server.CreateVirtualMachineResponseObject, error) {
	m, e := s.service.CreateVirtualMachine(r.Body.Name, r.Body.Cluster)

	if validation.Is(e) {
		return server.CreateVirtualMachine400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreateVirtualMachine503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateVirtualMachine500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.CreateVirtualMachine201JSONResponse(
			*convert.VirtualMachine(m),
		),
		nil
}
