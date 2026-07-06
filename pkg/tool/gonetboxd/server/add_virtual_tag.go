package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) AddVirtualTag(
	_ context.Context,
	r server.AddVirtualTagRequestObject,
) (server.AddVirtualTagResponseObject, error) {
	vm, e := s.client.AddVirtualTag(r.Name, r.Tag)

	if e != nil {
		return server.AddVirtualTag500JSONResponse(*s.captureDetail(e)), nil
	}

	entry := server.VirtualMachine{Identifier: vm.Identifier, Name: vm.Name}

	if len(vm.Tags) > 0 {
		entry.Tags = &vm.Tags
	}

	return server.AddVirtualTag200JSONResponse(entry), nil
}
