package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) DeleteMachine(
	_ context.Context,
	r server.DeleteMachineRequestObject,
) (server.DeleteMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.DeleteMachine400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.DeleteMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	purge := false

	if r.Params.Purge != nil {
		purge = *r.Params.Purge
	}

	e = s.service.DeleteMachine(c, int(r.Identifier), node, purge)

	if e != nil {
		if not_found.Is(e) {
			return server.DeleteMachine404JSONResponse{Error: e.Error()}, nil
		}

		if conflict.Is(e) {
			return server.DeleteMachine409JSONResponse(*clientError(e)), nil
		}

		return server.DeleteMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteMachine200JSONResponse{TaskId: "deleted"}, nil
}
