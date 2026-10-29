package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListVirtualMachines(
	_ context.Context,
	_ server.ListVirtualMachinesRequestObject,
) (server.ListVirtualMachinesResponseObject, error) {
	machines, grouped, e := s.service.VirtualMachines()

	if forbidden.Is(e) {
		return server.ListVirtualMachines403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListVirtualMachines500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListVirtualMachines200JSONResponse(
			convert.LabeledVirtualMachines(machines, grouped),
		),
		nil
}
