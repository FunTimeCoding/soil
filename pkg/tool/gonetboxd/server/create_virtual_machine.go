package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateVirtualMachine(
	_ context.Context,
	r server.CreateVirtualMachineRequestObject,
) (server.CreateVirtualMachineResponseObject, error) {
	m, e := s.service.CreateVirtualMachine(r.Body.Name, r.Body.Cluster)

	if validation.Is(e) {
		return server.CreateVirtualMachine400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.CreateVirtualMachine403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateVirtualMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateVirtualMachine201JSONResponse(
			*convert.VirtualMachine(m),
		),
		nil
}
