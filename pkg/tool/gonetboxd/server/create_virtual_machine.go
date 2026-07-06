package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateVirtualMachine(
	_ context.Context,
	r server.CreateVirtualMachineRequestObject,
) (server.CreateVirtualMachineResponseObject, error) {
	cl, e := s.client.ClusterByName(r.Body.Cluster)

	if e != nil {
		return server.CreateVirtualMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	vm, f := s.client.CreateVirtualMachine(r.Body.Name, cl)

	if f != nil {
		return server.CreateVirtualMachine500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.CreateVirtualMachine201JSONResponse{
		Identifier: vm.Identifier,
		Name:       vm.Name,
		Cluster:    &cl.Name,
	}, nil
}
