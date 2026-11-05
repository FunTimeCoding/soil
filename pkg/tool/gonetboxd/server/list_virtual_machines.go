package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListVirtualMachines(
	x context.Context,
	_ server.ListVirtualMachinesRequestObject,
) (server.ListVirtualMachinesResponseObject, error) {
	machines, grouped, e := s.service.VirtualMachines()

	if unavailable.Is(e) {
		return server.ListVirtualMachines503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListVirtualMachines500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListVirtualMachines200JSONResponse(
			convert.LabeledVirtualMachines(machines, grouped),
		),
		nil
}
