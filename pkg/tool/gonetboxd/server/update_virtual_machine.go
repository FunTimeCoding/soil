package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) UpdateVirtualMachine(
	_ context.Context,
	r server.UpdateVirtualMachineRequestObject,
) (server.UpdateVirtualMachineResponseObject, error) {
	if r.Body.PrimaryAddress != nil {
		if _, e := s.client.SetVirtualMachinePrimaryAddress(
			r.Name,
			*r.Body.PrimaryAddress,
		); e != nil {
			return server.UpdateVirtualMachine500JSONResponse(
				*s.captureDetail(e),
			), nil
		}
	}

	name := r.Name

	if r.Body.Name != nil {
		result, e := s.client.RenameVirtualMachine(r.Name, *r.Body.Name)

		if e != nil {
			return server.UpdateVirtualMachine500JSONResponse(
				*s.captureDetail(e),
			), nil
		}

		name = result.Name
	}

	m, e := s.client.VirtualMachineByName(name)

	if e != nil {
		return server.UpdateVirtualMachine500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.UpdateVirtualMachine200JSONResponse(
		*convert.VirtualMachine(m),
	), nil
}
