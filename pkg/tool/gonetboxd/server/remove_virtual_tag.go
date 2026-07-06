package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) RemoveVirtualTag(
	_ context.Context,
	r server.RemoveVirtualTagRequestObject,
) (server.RemoveVirtualTagResponseObject, error) {
	vm, e := s.client.RemoveVirtualTag(r.Name, r.Tag)

	if e != nil {
		return server.RemoveVirtualTag500JSONResponse(*s.captureDetail(e)), nil
	}

	entry := server.VirtualMachine{Identifier: vm.Identifier, Name: vm.Name}

	if len(vm.Tags) > 0 {
		entry.Tags = &vm.Tags
	}

	return server.RemoveVirtualTag200JSONResponse(entry), nil
}
