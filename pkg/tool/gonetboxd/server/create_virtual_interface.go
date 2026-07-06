package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateVirtualInterface(
	_ context.Context,
	r server.CreateVirtualInterfaceRequestObject,
) (server.CreateVirtualInterfaceResponseObject, error) {
	vm, e := s.client.VirtualMachineByName(r.Name)

	if e != nil {
		return server.CreateVirtualInterface500JSONResponse(*s.captureDetail(e)), nil
	}

	i, f := s.client.CreateVirtualInterface(vm, r.Body.Name)

	if f != nil {
		return server.CreateVirtualInterface500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.CreateVirtualInterface201JSONResponse{
		Identifier: i.GetId(), Name: i.GetName(),
	}, nil
}
