package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) AddVirtualTag(
	_ context.Context,
	r server.AddVirtualTagRequestObject,
) (server.AddVirtualTagResponseObject, error) {
	vm, e := s.service.AddVirtualTag(r.Name, r.Tag)

	if not_found.Is(e) {
		return server.AddVirtualTag404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.AddVirtualTag400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.AddVirtualTag403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.AddVirtualTag500JSONResponse(*s.captureDetail(e)), nil
	}

	entry := server.VirtualMachine{Identifier: vm.Identifier, Name: vm.Name}

	if len(vm.Tags) > 0 {
		entry.Tags = &vm.Tags
	}

	return server.AddVirtualTag200JSONResponse(entry), nil
}
