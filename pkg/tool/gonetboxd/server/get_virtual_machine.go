package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) GetVirtualMachine(
	_ context.Context,
	r server.GetVirtualMachineRequestObject,
) (server.GetVirtualMachineResponseObject, error) {
	m, labels, e := s.service.VirtualMachine(r.Name)

	if not_found.Is(e) {
		return server.GetVirtualMachine404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.GetVirtualMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.GetVirtualMachine200JSONResponse(
		*convert.LabeledVirtualMachine(m, labels),
	), nil
}
