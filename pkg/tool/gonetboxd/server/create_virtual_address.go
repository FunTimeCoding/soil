package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateVirtualAddress(
	_ context.Context,
	r server.CreateVirtualAddressRequestObject,
) (server.CreateVirtualAddressResponseObject, error) {
	vm, e := s.client.VirtualMachineByName(r.Name)

	if e != nil {
		return server.CreateVirtualAddress500JSONResponse(*s.captureDetail(e)), nil
	}

	i, f := s.client.VirtualMachineInterfaceByName(vm, r.Body.Interface)

	if f != nil {
		return server.CreateVirtualAddress500JSONResponse(*s.captureDetail(f)), nil
	}

	a, g := s.client.CreateVirtualAddress(i.GetId(), r.Body.Address)

	if g != nil {
		return server.CreateVirtualAddress500JSONResponse(*s.captureDetail(g)), nil
	}

	return server.CreateVirtualAddress201JSONResponse{
		Identifier: a.Identifier, Address: a.Address.String(),
	}, nil
}
