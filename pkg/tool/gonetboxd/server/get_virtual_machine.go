package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) GetVirtualMachine(
	_ context.Context,
	r server.GetVirtualMachineRequestObject,
) (server.GetVirtualMachineResponseObject, error) {
	m, e := s.client.VirtualMachineByName(r.Name)

	if e != nil {
		return server.GetVirtualMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	labels, f := s.store.Labels(constant.VirtualMachineAddress, m.Identifier)

	if f != nil {
		return server.GetVirtualMachine500JSONResponse(*s.captureDetail(f)), nil
	}

	result := convert.VirtualMachine(m)
	result.Labels = new(convert.Labels(labels))

	return server.GetVirtualMachine200JSONResponse(*result), nil
}
